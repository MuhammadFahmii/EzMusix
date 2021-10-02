package tracks

type TracksUsecase struct {
	trackRepository Repository
}

func NewTracksUsecase(repo Repository) Usecase {
	return &TracksUsecase{
		trackRepository: repo,
	}
}

func (trackUsecase *TracksUsecase) Get(track Domain) (Domain, error) {
	res, err := trackUsecase.trackRepository.Get(track)
	if err != nil {
		return Domain{}, nil
	}
	return res, nil
}
