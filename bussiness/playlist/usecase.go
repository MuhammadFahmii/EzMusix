package playlist

type PlaylistUsecase struct {
	playlistRepo Repository
}

func NewPlaylistUsecase(repo Repository) Usecase {
	return &PlaylistUsecase{
		playlistRepo: repo,
	}
}

func (playlistUsecase *PlaylistUsecase) Get(playlist *Domain) ([]Domain, error) {
	res, err := playlistUsecase.playlistRepo.Get(playlist)
	if err != nil {
		return []Domain{}, nil
	}
	return res, nil
}
