package playlist

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/tracks"
	"EzMusix/repository/thirdparty"
)

type Playlist struct {
	Id     int
	Name   string
	UserID int
	Tracks []thirdparty.Track `gorm:"many2many:detail_playlist"`
}

func fromDomain(domain playlists.Domain) Playlist {
	return Playlist{
		Id:     domain.Id,
		Name:   domain.Name,
		UserID: domain.UserID,
	}
}

func (pl *Playlist) toDomain() playlists.Domain {
	return playlists.Domain{
		Id:     pl.Id,
		Name:   pl.Name,
		UserID: pl.UserID,
		Tracks: toTrackDomain(pl),
	}
}
func toTrackDomain(pl *Playlist) []tracks.Domain {
	track := playlists.Domain{}.Tracks
	for _, val := range pl.Tracks {
		track = append(track, tracks.Domain(val))
	}
	return track
}
