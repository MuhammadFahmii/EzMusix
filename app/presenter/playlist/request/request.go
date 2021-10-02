package request

import "EzMusix/bussiness/playlist"

type Playlist struct {
	Name   string `json:"name" form:"name"`
	UserID int    `json:"user_id" form:"user_id"`
}

type DeletePlaylist struct {
	Id int `param:"id"`
}

func ToDomain(pl Playlist) playlist.Domain {
	return playlist.Domain{
		Name:   pl.Name,
		UserID: pl.UserID,
	}
}
func DeleteToDomain(pl DeletePlaylist) playlist.Domain {
	return playlist.Domain{
		Id: pl.Id,
	}
}
