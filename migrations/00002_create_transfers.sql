-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    transfers (
        id SERIAL PRIMARY KEY,
        counterparty_name TEXT NOT NULL,
        counterparty_iban TEXT NOT NULL,
        counterparty_bic TEXT NOT NULL,
        amount_cents INT NOT NULL,
        bank_account_id INT NOT NULL,
        description TEXT NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS
    transfers;
-- +goose StatementEnd
