package main

import (
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
	e.GET("/", handleGetAccount)

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

func handleGetAccount(c echo.Context) error {
	// TODO: Implement handleGetAccount function
	return c.String(http.StatusOK, "handle get account function!")
}
