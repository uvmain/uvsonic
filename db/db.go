package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(path string) {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("DB init failed:", err)
	}
	createTables()
}

func createTables() {
	schema := `
	CREATE TABLE IF NOT EXISTS artists (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS albums (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		artist_id TEXT,
		FOREIGN KEY (artist_id) REFERENCES artists(id)
	);
	CREATE TABLE IF NOT EXISTS songs (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		duration INTEGER NOT NULL,
		artist_id TEXT,
		album_id TEXT,
		FOREIGN KEY (artist_id) REFERENCES artists(id),
		FOREIGN KEY (album_id) REFERENCES albums(id)
	);
	`
	if _, err := DB.Exec(schema); err != nil {
		log.Fatal("Schema migration failed:", err)
	}
}
