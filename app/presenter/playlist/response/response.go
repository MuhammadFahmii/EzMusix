package response

import (
	"EzMusix/bussiness/playlist"
	"EzMusix/bussiness/tracks"
)

type Playlist struct {
	Id     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	UserID int    `json:"user_id" form:"user_id"`
	Track  []tracks.Track
}

type DeletePlaylist struct {
	Name string
}

func FromDomain(pl playlist.Playlist) *Playlist {
	return &Playlist{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
		Track:  pl.Tracks,
	}
}
