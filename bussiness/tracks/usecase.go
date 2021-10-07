package tracks

import (
	"errors"
)

type TracksUsecase struct {
	trackRepository Repository
	thirdParty      ThirdParty
}

func NewTracksUsecase(tracksRepo Repository, thirdParty ThirdParty) Usecase {
	return &TracksUsecase{
		trackRepository: tracksRepo,
		thirdParty:      thirdParty,
	}
}

func (trackUsecase *TracksUsecase) Get(trackName, artistName string) (Domain, error) {
	res, err := trackUsecase.thirdParty.Get(trackName, artistName)
	if res.Id == 0 {
		return Domain{}, errors.New("not found")
	}
	if err != nil {
		return Domain{}, nil
	}
	return res, nil
}
func (tracksUsecase *TracksUsecase) AddTrackPlaylist(detailPlaylist TrackPlaylist) (Domain, error) {
	res, err := tracksUsecase.thirdParty.Get(detailPlaylist.TrackName, detailPlaylist.ArtistName)
	if err != nil {
		return Domain{}, nil
	}
	newTrack, err := tracksUsecase.trackRepository.AddTrackPlaylist(detailPlaylist, res)
	if err != nil {
		return Domain{}, err
	}
	return newTrack, nil
}

func (tracksUsecase *TracksUsecase) DeleteTrackPlaylist(playlistId, trackId int) (DeleteTrackPlaylist, error) {
	if playlistId == 0 || trackId == 0 {
		return DeleteTrackPlaylist{}, errors.New("please fill all param")
	}
	deleteTrackPlaylist, err := tracksUsecase.trackRepository.DeleteTrackPlaylist(playlistId, trackId)
	if err != nil {
		return DeleteTrackPlaylist{}, err
	}
	return deleteTrackPlaylist, nil
}
