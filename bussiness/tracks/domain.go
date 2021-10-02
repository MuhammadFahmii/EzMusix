package tracks

type Domain struct {
	Id         int
	Name       string
	ArtistName string
	AlbumName  string
}

type Usecase interface {
	Get(track Domain) (Domain, error)
}

type Repository interface {
	Get(track Domain) (Domain, error)
}
