package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/uvmain/uvsonic/logic"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	logic.CreateDir(logic.DatabaseDirectory)

	dbPath := filepath.Join(logic.DatabaseDirectory, "sqlite.db")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("Creating database file")

		file, err := os.Create(dbPath)
		if err != nil {
			log.Printf("Error creating database file: %s", err)
		} else {
			log.Println("Database file created")
		}
		file.Close()
	} else {
		log.Println("Database already exists")
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("DB init failed:", err)
	}

	_, err = DB.Exec("pragma journal_mode = wal;")
	if err != nil {
		log.Printf("Error entering WAL mode: %s", err)
	} else {
		log.Println("Database is in WAL mode")
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
