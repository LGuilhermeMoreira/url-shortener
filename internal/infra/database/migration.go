package database

import "database/sql"

func Migration(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS urls (
    	short_id VARCHAR(10) NOT NULL UNIQUE,
    	url TEXT NOT NULL
	);`)
	return err
}
