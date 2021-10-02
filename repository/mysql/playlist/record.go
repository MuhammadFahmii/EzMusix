package playlist

import (
	"EzMusix/bussiness/playlist"
	"EzMusix/repository/mysql/tracks"
)

type Playlist struct {
	Id     int
	Name   string
	UserID int
	Tracks []tracks.Track `gorm:"many2many:detail_playlist"`
}

func FromDomain(domain *playlist.Playlist) Playlist {
	return Playlist{
		Id:     domain.Id,
		Name:   domain.Name,
		UserID: domain.UserID,
	}
}

func ToDomain(pl *Playlist) playlist.Playlist {
	return playlist.Playlist{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
		Tracks: pl.Tracks,
	}
}
