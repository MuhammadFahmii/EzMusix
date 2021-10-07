package tracks

import "time"

type Domain struct {
	Id            int
	TrackName     string
	ArtistName    string
	AlbumName     string
	TrackRating   int
	TrackShareUrl string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TrackPlaylist struct {
	PlaylistId int
	ArtistName string
	TrackName  string
}

type DeleteTrackPlaylist struct {
	PlaylistName string
	TrackName    string
}

type Usecase interface {
	Get(trackName, artistName string) (Domain, error)
	AddTrackPlaylist(detailPlaylist TrackPlaylist) (Domain, error)
	DeleteTrackPlaylist(playlistId, trackId int) (DeleteTrackPlaylist, error)
}

type Repository interface {
	AddTrackPlaylist(TrackPlaylist, Domain) (Domain, error)
	DeleteTrackPlaylist(playlistId, trackId int) (DeleteTrackPlaylist, error)
}

type ThirdParty interface {
	Get(trackName, artistName string) (Domain, error)
}
