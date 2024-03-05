-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS doctors (
    doctor_id SERIAL PRIMARY KEY,
    specialties VARCHAR(255) NOT NULL,
    user_id SERIAL,

    FOREIGN KEY (user_id) REFERENCES users(user_id) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctors;
-- +goose StatementEnd
