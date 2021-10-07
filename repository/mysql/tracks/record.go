package tracks

import (
	"EzMusix/bussiness/tracks"
	"time"
)

type Track struct {
	Id            int
	TrackName     string
	ArtistName    string
	AlbumName     string
	TrackRating   int
	TrackShareUrl string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Playlist struct {
	Id        int
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Tracks    []Track `gorm:"many2many:detail_playlist"`
}

func FromDomain(tracks tracks.Domain) Track {
	return Track{
		Id:            tracks.Id,
		TrackName:     tracks.TrackName,
		ArtistName:    tracks.ArtistName,
		AlbumName:     tracks.AlbumName,
		TrackRating:   tracks.TrackRating,
		TrackShareUrl: tracks.TrackShareUrl,
		CreatedAt:     tracks.CreatedAt,
		UpdatedAt:     tracks.UpdatedAt,
	}
}

func ToDomain(track Track) tracks.Domain {
	return tracks.Domain{
		Id:            track.Id,
		TrackName:     track.TrackName,
		ArtistName:    track.ArtistName,
		AlbumName:     track.AlbumName,
		TrackRating:   track.TrackRating,
		TrackShareUrl: track.TrackShareUrl,
		CreatedAt:     track.CreatedAt,
		UpdatedAt:     track.UpdatedAt,
	}
}
