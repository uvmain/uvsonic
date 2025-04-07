package models

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ArtistID string `json:"artist_id"`
}

type Song struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	ArtistID string `json:"artist_id"`
	AlbumID  string `json:"album_id"`
}
