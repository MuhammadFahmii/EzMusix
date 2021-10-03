package tracks

type TracksUsecase struct {
	trackRepository ThirdParty
}

func NewTracksUsecase(repo ThirdParty) Usecase {
	return &TracksUsecase{
		trackRepository: repo,
	}
}

func (trackUsecase *TracksUsecase) Get(trackName, artistName string) (Track, error) {
	res, err := trackUsecase.trackRepository.Get(trackName, artistName)
	if err != nil {
		return Track{}, nil
	}
	return res, nil
}
func (TracksUsecase *TracksUsecase) AddDetailPlaylist(detailPlaylist DetailPlaylist) (Track, error) {
	newTrack, err := TracksUsecase.trackRepository.AddDetailPlaylist(detailPlaylist)
	if err != nil {
		return Track{}, err
	}
	return newTrack, nil
}

func (TracksUsecase *TracksUsecase) DeleteDetailPlaylist(playlistId, trackId int) (DetailPlaylist, error) {
	deletePlaylistTrack, err := TracksUsecase.trackRepository.DeleteDetailPlaylist(playlistId, trackId)
	if err != nil {
		return DetailPlaylist{}, err
	}
	return deletePlaylistTrack, nil
}
