package playlists

import "errors"

type PlaylistUsecase struct {
	playlistRepo Repository
}

func NewPlaylistUsecase(repo Repository) Usecase {
	return &PlaylistUsecase{
		playlistRepo: repo,
	}
}

func (playlistUseCase *PlaylistUsecase) Insert(playlist Domain) (Domain, error) {
	res, err := playlistUseCase.playlistRepo.Insert(playlist)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
func (playlistUseCase *PlaylistUsecase) Delete(playlist Domain) (Domain, error) {
	res, err := playlistUseCase.playlistRepo.Delete(playlist)
	if err != nil {
		if err.Error() == "record not found" {
			return Domain{}, errors.New("record not found")
		}
		return Domain{}, err
	}
	return res, nil
}

func (playlistUseCase *PlaylistUsecase) Get(playlist Domain) ([]Domain, error) {
	res, err := playlistUseCase.playlistRepo.Get(playlist)
	if err != nil {
		if err.Error() == "record not found" {
			return []Domain{}, errors.New("not found")
		}
		return []Domain{}, err
	}
	return res, nil
}
