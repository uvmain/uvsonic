package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/uvmain/uvsonic/db"
	"github.com/uvmain/uvsonic/models"
)

func HandleSongs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := db.DB.Query("SELECT id, title, duration, artist_id, album_id FROM songs")
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var albums []models.Song
		for rows.Next() {
			var a models.Song
			if err := rows.Scan(&a.ID, &a.Title, &a.Duration, &a.ArtistID, &a.AlbumID); err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				return
			}
			albums = append(albums, a)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(albums)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
