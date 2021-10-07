package playlist

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/tracks"
	trackRepo "EzMusix/repository/mysql/tracks"
	"time"
)

type Playlist struct {
	Id        int
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Tracks    []trackRepo.Track `gorm:"many2many:detail_playlist"`
}

func fromDomain(domain playlists.Domain) Playlist {
	return Playlist{
		Id:        domain.Id,
		Name:      domain.Name,
		UserID:    domain.UserID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (pl *Playlist) toDomain() playlists.Domain {
	return playlists.Domain{
		Id:        pl.Id,
		Name:      pl.Name,
		UserID:    pl.UserID,
		CreatedAt: pl.CreatedAt,
		UpdatedAt: pl.UpdatedAt,
		Tracks:    convertToArray(pl.Tracks),
	}
}

func convertToArray(track []trackRepo.Track) []tracks.Domain {
	tracksDomain := []tracks.Domain{}
	for _, val := range track {
		tracksDomain = append(tracksDomain, toTrackDomain(val))
	}
	return tracksDomain
}

func toTrackDomain(track trackRepo.Track) tracks.Domain {
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
