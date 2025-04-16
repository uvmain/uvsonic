package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/uvmain/uvsonic/logic"
	"github.com/uvmain/uvsonic/types"
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
	err := CreateTrackMetadataTable()
	if err != nil {
		log.Fatal("Failed to create TrackMetadataTable:", err)
	}
}

func CreateTrackMetadataTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS track_metadata (
		musicbrainz_track_id TEXT PRIMARY KEY,
		filename TEXT,
		format TEXT,
		duration TEXT,
		size TEXT,
		bitrate TEXT,
		title TEXT,
		artist TEXT,
		album TEXT,
		album_artist TEXT,
		genre TEXT,
		track_number TEXT,
		total_tracks TEXT,
		disc_number TEXT,
		total_discs TEXT,
		release_date TEXT,
		musicbrainz_artist_id TEXT,
		musicbrainz_album_id TEXT,
		label TEXT
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("Error creating track_metadata table: %v", err)
		return err
	}

	log.Println("track_metadata table created successfully (if it didn't already exist).")
	return nil
}

func InsertTrackMetadata(metadata types.TrackMetadata) error {
	stmt, err := DB.Prepare(`INSERT INTO track_metadata (
		musicbrainz_track_id, filename, format, duration, size, bitrate, title, artist, album,
		album_artist, genre, track_number, total_tracks, disc_number, total_discs, release_date,
		musicbrainz_artist_id, musicbrainz_album_id, label 
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		metadata.MusicBrainzTrackID, metadata.Filename, metadata.Format, metadata.Duration, metadata.Size,
		metadata.Bitrate, metadata.Title, metadata.Artist, metadata.Album,
		metadata.AlbumArtist, metadata.Genre, metadata.TrackNumber,
		metadata.TotalTracks, metadata.DiscNumber, metadata.TotalDiscs,
		metadata.ReleaseDate, metadata.MusicBrainzArtistID,
		metadata.MusicBrainzAlbumID, metadata.Label,
	)
	if err != nil {
		log.Printf("error inserting metadata row: %s", err)
		return err
	}

	log.Printf("Metadata row inserted successfully for %s", metadata.Filename)
	return nil
}
