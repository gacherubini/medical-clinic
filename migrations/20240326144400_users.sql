-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD hash_password VARCHAR(255) NOT NULL,
ADD email VARCHAR(255) NOT NULL

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS hash_password,
DROP COLUMN IF EXISTS email;
-- +goose StatementEnd
