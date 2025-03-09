package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"eggs.corm/eggs/bts/migrations"
	"eggs.corm/eggs/bts/models"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type config struct {
	PostgresConnectionString string
	Server                   struct {
		Port string
	}
}

type Transaction struct {
	Amount           string `json:"amount"`
	CounterpartyName string `json:"counterparty_name"`
	CounterpartyBic  string `json:"counterparty_bic"`
	CounterpartyIban string `json:"counterparty_iban"`
	Description      string `json:"description"`
}

type BulkTransaction struct {
	OrganizationName string        `json:"organization_name"`
	OrganizationBic  string        `json:"organization_bic"`
	OrganizationIban string        `json:"organization_iban"`
	CreditTransfers  []Transaction `json:"credit_transfers"`
}

func main() {
	// setup configuration
	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	db, err := models.OpenDbConnection(cfg.PostgresConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.ApplyDatabaseMigrations(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// setup server
	e := echo.New()
	e.POST("/salaries", handlePostAccount)

	// start server
	e.Logger.Fatal(e.Start(cfg.Server.Port))
}

func loadEnvConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	cfg.PostgresConnectionString = os.Getenv("BTS_DB_CONNECTION_STRING")
	cfg.Server.Port = os.Getenv("BTS_SERVER_PORT")

	return cfg, nil
}

func handlePostAccount(c echo.Context) error {
	err := validateRequest(c)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity,
			err.Error())
	}

	bt := new(BulkTransaction)
	if err = c.Bind(bt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func validateRequest(c echo.Context) error {
	if v := c.Request().ContentLength <= 0; v {
		errMsg := fmt.Sprintf("Request was empty. Content length = %v\n",
			c.Request().ContentLength)
		fmt.Printf("%s\n", errMsg)
		return errors.New(errMsg)
	}
	return nil
}

// TODO: Implement bulk transfer handler
// 3. In successfull case add transfet to db, update customer's balance and
// 	  return 201 http status code.
