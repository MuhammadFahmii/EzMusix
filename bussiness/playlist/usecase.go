package playlist

type PlaylistUsecase struct {
	playlistRepo Repository
}

func NewPlaylistUsecase(repo Repository) Usecase {
	return &PlaylistUsecase{
		playlistRepo: repo,
	}
}

func (playlistUseCase *PlaylistUsecase) Insert(playlist Playlist) (Playlist, error) {
	res, err := playlistUseCase.playlistRepo.Insert(playlist)
	if err != nil {
		return Playlist{}, nil
	}
	return res, nil
}
func (playlistUseCase *PlaylistUsecase) Delete(playlist Playlist) (Playlist, error) {
	res, err := playlistUseCase.playlistRepo.Delete(playlist)
	if err != nil {
		return Playlist{}, nil
	}
	return res, nil
}

func (playlistUseCase *PlaylistUsecase) Get(playlist Playlist) ([]Playlist, error) {
	res, err := playlistUseCase.playlistRepo.Get(playlist)
	if err != nil {
		return []Playlist{}, nil
	}
	return res, nil
}
