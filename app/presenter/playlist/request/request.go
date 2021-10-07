package request

import (
	"EzMusix/bussiness/playlists"
	"time"
)

type Playlist struct {
	Name      string    `json:"name" form:"name"`
	UserID    int       `json:"user_id" form:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePlaylist struct {
	Id int `param:"id"`
}

func ToDomain(pl Playlist) playlists.Domain {
	return playlists.Domain{
		Name:   pl.Name,
		UserID: pl.UserID,
	}
}
func DeleteToDomain(pl DeletePlaylist) playlists.Domain {
	return playlists.Domain{
		Id: pl.Id,
	}
}
