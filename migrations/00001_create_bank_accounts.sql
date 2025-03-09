-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    bank_accounts (
        id SERIAL PRIMARY KEY,
        organization_name TEXT NOT NULL,
        balance_cents INT NOT NULL,
        iban TEXT NOT NULL,
        bic TEXT NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS
    bank_accounts;
-- +goose StatementEnd