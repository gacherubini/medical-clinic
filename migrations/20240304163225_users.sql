-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
