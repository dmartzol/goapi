package postgres

import (
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

const (
	migrationTable = "goose_migrations"
	migrationPath  = "internal/storage/postgres/sql"
)

func (db *DB) Migrate() error {
	goose.SetTableName(migrationTable)
	err := goose.Up(db.Client.DB, migrationPath)
	if err != nil {
		return errors.Wrap(err, "failed migration")
	}

	return nil
}
