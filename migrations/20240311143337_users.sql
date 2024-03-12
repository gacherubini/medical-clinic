-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD role varchar(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
