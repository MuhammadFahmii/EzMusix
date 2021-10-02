package request

import "EzMusix/bussiness/playlist"

type Playlist struct {
	Id     int `json:"id"`
	UserID int `json:"user_id"`
}

func ToDomain(pl Playlist) *playlist.Domain {
	return &playlist.Domain{
		Id:     pl.Id,
		UserID: pl.UserID,
	}
}
