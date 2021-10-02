package request

import "EzMusix/bussiness/playlist"

type Playlist struct {
	Id     int `json:"id"`
	UserID int `json:"user_id"`
}

func ToDomain(pl Playlist) *playlist.Playlist {
	return &playlist.Playlist{
		Id:     pl.Id,
		UserID: pl.UserID,
	}
}
