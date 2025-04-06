package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/uvmain/uvsonic/db"
	"github.com/uvmain/uvsonic/models"
)

func HandleAlbums(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := db.DB.Query("SELECT id, name, artist_id FROM albums")
		if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var albums []models.Album
		for rows.Next() {
			var a models.Album
			if err := rows.Scan(&a.ID, &a.Name, &a.ArtistID); err != nil {
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
