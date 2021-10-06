package tracks

type Track struct {
	Id         int
	Name       string
	ArtistName string
	AlbumName  string
}

type DetailPlaylist struct {
	PlaylistId int
	ArtistName string
	TrackName  string
}

type Usecase interface {
	Get(trackName, artistName string) (Track, error)
	AddDetailPlaylist(detailPlaylist DetailPlaylist) (Track, error)
	DeleteDetailPlaylist(playlistId, trackId int) (DetailPlaylist, error)
}

type ThirdParty interface {
	Get(trackName, artistName string) (Track, error)
	AddDetailPlaylist(detailPlaylist DetailPlaylist) (Track, error)
	DeleteDetailPlaylist(playlistId, trackId int) (DetailPlaylist, error)
}
