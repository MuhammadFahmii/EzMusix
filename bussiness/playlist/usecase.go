package playlist

type PlaylistUsecase struct {
	playlistRepo Repository
}

func NewPlaylistUsecase(repo Repository) Usecase {
	return &PlaylistUsecase{
		playlistRepo: repo,
	}
}

func (playlistUsecase *PlaylistUsecase) Insert(playlist Domain) (Domain, error) {
	res, err := playlistUsecase.playlistRepo.Insert(playlist)
	if err != nil {
		return Domain{}, nil
	}
	return res, nil
}

func (playlistUsecase *PlaylistUsecase) Get(playlist Domain) ([]Domain, error) {
	res, err := playlistUsecase.playlistRepo.Get(playlist)
	if err != nil {
		return []Domain{}, nil
	}
	return res, nil
}

func (playlistUsecase *PlaylistUsecase) Delete(playlist Domain) (Domain, error) {
	res, err := playlistUsecase.playlistRepo.Delete(playlist)
	if err != nil {
		return Domain{}, nil
	}
	return res, nil
}
