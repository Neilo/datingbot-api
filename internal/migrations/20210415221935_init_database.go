package migrations

import (
	"database/sql"
	"errors"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInitDatabase, downInitDatabase)
}

func upInitDatabase(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE bot_users (
    id TEXT PRIMARY KEY,
    chat_id TEXT,
    name TEXT,
    description TEXT,
    age BIGINT,
    gender TEXT,
    find_sex TEXT,
    longitude decimal,
    latitude decimal,
    rating decimal
);
`)
	return err
}

func downInitDatabase(tx *sql.Tx) error {
	return errors.New("1rst version of database")
}
