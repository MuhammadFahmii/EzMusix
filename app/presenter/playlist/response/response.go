package response

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/tracks"
)

type Playlist struct {
	Id     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	UserID int    `json:"user_id" form:"user_id"`
	Track  []tracks.Domain
}

type DeletePlaylist struct {
	Name string
}

func FromDomain(pl playlists.Domain) *Playlist {
	return &Playlist{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
		Track:  pl.Tracks,
	}
}
