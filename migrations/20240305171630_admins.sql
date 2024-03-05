-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS admins (
    admin_id SERIAL PRIMARY KEY,
    user_id SERIAL,

    FOREIGN KEY (user_id) REFERENCES users(user_id) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS admins;
-- +goose StatementEnd
