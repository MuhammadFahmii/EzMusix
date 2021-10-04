package tracks

type Domain struct {
	Id         int
	Name       string
	ArtistName string
	AlbumName  string
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
	AddDetailPlaylist(detailPlaylist TrackPlaylist) (Domain, error)
	DeleteDetailPlaylist(playlistId, trackId int) (DeleteTrackPlaylist, error)
}

type ThirdParty interface {
	Get(trackName, artistName string) (Domain, error)
	AddDetailPlaylist(detailPlaylist TrackPlaylist) (Domain, error)
	DeleteDetailPlaylist(playlistId, trackId int) (DeleteTrackPlaylist, error)
}
