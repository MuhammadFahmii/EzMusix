package playlist

import "EzMusix/bussiness/tracks"

type Domain struct {
	Id     int
	Name   string
	UserID int
	Tracks []tracks.Domain
}

type Usecase interface {
	Get(playlist *Domain) ([]Domain, error)
}

type Repository interface {
	Get(playlist *Domain) ([]Domain, error)
}
