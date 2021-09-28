package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddLike, downAddLike)
}

func upAddLike(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE bot_likes (
    id TEXT PRIMARY KEY,
    subject TEXT,
    object TEXT,
    create_at BIGINT
);
`)
	return err
}

func downAddLike(tx *sql.Tx) error {
	_, err := tx.Exec(`
	DROP table bot_likes; 
`)
	return err
}
