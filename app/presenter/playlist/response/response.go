package response

import "EzMusix/bussiness/playlist"

type Playlist struct {
	Id     int    `json:"id"`
	Name   string `json:"name" form:"name"`
	UserID int    `json:"user_id" form:"user_id"`
}

type DeletePlaylist struct {
	Name string
}

func FromDomain(pl playlist.Domain) *Playlist {
	return &Playlist{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
	}
}
