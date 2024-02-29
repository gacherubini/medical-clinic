package main

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpEx2, DownEx2)
}

func UpEx2(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS universities (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        country VARCHAR(255) NOT NULL);`)

	if err != nil {
		return err
	}

	return nil
}

func DownEx2(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS universities;`)
	if err != nil {
		return err
	}
	return nil
}
