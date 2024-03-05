-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS patients (
    patient_id SERIAL PRIMARY KEY,
    user_id SERIAL, 

    FOREIGN KEY (user_id) REFERENCES users(user_id) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS patients;
-- +goose StatementEnd
