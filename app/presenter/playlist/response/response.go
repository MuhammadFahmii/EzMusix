package response

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/tracks"
)

type Playlist struct {
	Id    int    `json:"id"`
	Name  string `json:"name" form:"name"`
	Track []tracks.Domain
}

type AddPlaylist struct {
	Name string `json:"name"`
}

type DeletePlaylist struct {
	Name string
}

func FromDomain(pl playlists.Domain) Playlist {
	return Playlist{
		Id:    pl.Id,
		Name:  pl.Name,
		Track: pl.Tracks,
	}
}

func ToAddPlaylist(pl playlists.Domain) AddPlaylist {
	return AddPlaylist{
		Name: pl.Name,
	}
}
