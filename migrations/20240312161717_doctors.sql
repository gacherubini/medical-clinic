-- +goose Up
-- +goose StatementBegin
ALTER TABLE doctors
ADD healthinsurance_id INT,
ADD CONSTRAINT fk_healthinsurance_id FOREIGN KEY (healthinsurance_id) REFERENCES healthinsurance(healthinsurance_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE doctors
DROP COLUMN IF EXISTS healthinsurance_id;
-- +goose StatementEnd

