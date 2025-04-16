package types

type TrackMetadata struct {
	Filename            string `json:"filename"`
	Format              string `json:"format"`
	Duration            string `json:"duration"`
	Size                string `json:"size"`
	Bitrate             string `json:"bitrate"`
	Title               string `json:"title"`
	Artist              string `json:"artist"`
	Album               string `json:"album"`
	AlbumArtist         string `json:"album_artist"`
	Genre               string `json:"genre"`
	TrackNumber         string `json:"track_number"`
	DiscNumber          string `json:"disc_number"`
	ReleaseDate         string `json:"release_date"`
	MusicBrainzArtistID string `json:"musicbrainz_artist_id"`
	MusicBrainzAlbumID  string `json:"musicbrainz_album_id"`
	Label               string `json:"label"`
}
