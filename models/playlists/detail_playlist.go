package playlists

type DetailPlaylist struct {
	PlaylistId int    `json:"playlist_id" form:"playlist_id"`
	TrackName  string `json:"track_name" form:"track_name"`
	ArtistName string `json:"artist_name" form:"artist_name"`
}
