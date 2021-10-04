package tracks

type TracksUsecase struct {
	trackRepository ThirdParty
}

func NewTracksUsecase(repo ThirdParty) Usecase {
	return &TracksUsecase{
		trackRepository: repo,
	}
}

func (trackUsecase *TracksUsecase) Get(trackName, artistName string) (Domain, error) {
	res, err := trackUsecase.trackRepository.Get(trackName, artistName)
	if err != nil {
		return Domain{}, nil
	}
	return res, nil
}
func (TracksUsecase *TracksUsecase) AddDetailPlaylist(detailPlaylist TrackPlaylist) (Domain, error) {
	newTrack, err := TracksUsecase.trackRepository.AddDetailPlaylist(detailPlaylist)
	if err != nil {
		return Domain{}, err
	}
	return newTrack, nil
}

func (TracksUsecase *TracksUsecase) DeleteDetailPlaylist(playlistId, trackId int) (DeleteTrackPlaylist, error) {
	deleteTrackPlaylist, err := TracksUsecase.trackRepository.DeleteDetailPlaylist(playlistId, trackId)
	if err != nil {
		return DeleteTrackPlaylist{}, err
	}
	return deleteTrackPlaylist, nil
}
