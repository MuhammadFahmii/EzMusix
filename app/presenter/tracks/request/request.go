package request

import "EzMusix/bussiness/tracks"

type Tracks struct {
	TrackName  string `json:"track_name" query:"track_name"`
	ArtistName string `json:"artist_name" query:"artist_name"`
}

func (rec Tracks) ToDomain() tracks.Domain {
	return tracks.Domain{
		Name:       rec.TrackName,
		ArtistName: rec.ArtistName,
	}
}
