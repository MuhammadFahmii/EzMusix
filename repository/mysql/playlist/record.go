package playlist

import (
	"EzMusix/bussiness/playlist"
	"EzMusix/bussiness/tracks"
)

type Playlist struct {
	Id     int
	Name   string
	UserID int
	Tracks []tracks.Track `gorm:"many2many:detail_playlist"`
}

func fromDomain(domain playlist.Playlist) Playlist {
	return Playlist{
		Id:     domain.Id,
		Name:   domain.Name,
		UserID: domain.UserID,
	}
}

func (pl *Playlist) toDomain() playlist.Playlist {
	return playlist.Playlist{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
		Tracks: pl.Tracks,
	}
}
