package playlist

import (
	"EzMusix/bussiness/playlist"
	"EzMusix/repository/thirdparty"
)

type Playlist struct {
	Id     int
	Name   string
	UserID int
	Tracks []thirdparty.Track `gorm:"many2many:detail_playlist"`
}

func fromDomain(domain playlist.Domain) Playlist {
	return Playlist{
		Id:     domain.Id,
		Name:   domain.Name,
		UserID: domain.UserID,
	}
}

func (pl *Playlist) toDomain() playlist.Domain {
	return playlist.Domain{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
		Tracks: pl.Tracks,
	}
}
