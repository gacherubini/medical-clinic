-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD role varchar(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS role;
-- +goose StatementEnd
