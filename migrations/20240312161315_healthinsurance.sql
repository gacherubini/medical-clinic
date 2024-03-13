-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS healthinsurance (
    healthinsurance_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS healthinsurance;
-- +goose StatementEnd
