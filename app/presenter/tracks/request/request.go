package request

import "EzMusix/bussiness/tracks"

type Tracks struct {
	TrackName  string `json:"track_name" query:"track_name"`
	ArtistName string `json:"artist_name" query:"artist_name"`
}

type DetailPlaylist struct {
	PlaylistId int    `json:"playlist_id" form:"playlist_id"`
	TrackName  string `json:"track_name" form:"track_name"`
	ArtistName string `json:"artist_name" form:"artist_name"`
}

type DeleteDetailPlaylist struct {
	PlaylistId int `json:"playlist_id" param:"playlist_id"`
	TrackId    int `json:"track_id" param:"track_id"`
}

func ToDetailPlaylist(dp DetailPlaylist) tracks.DetailPlaylist {
	return tracks.DetailPlaylist{
		PlaylistId: dp.PlaylistId,
		TrackName:  dp.TrackName,
		ArtistName: dp.ArtistName,
	}
}
