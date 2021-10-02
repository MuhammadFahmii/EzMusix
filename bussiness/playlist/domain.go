package playlist

import (
	"EzMusix/repository/thirdparty"
)

type Domain struct {
	Id     int
	Name   string
	UserID int
	Tracks []thirdparty.Track
}

type Usecase interface {
	Insert(playlist Domain) (Domain, error)
	Get(playlist Domain) ([]Domain, error)
	Delete(playlist Domain) (Domain, error)
}

type Repository interface {
	Insert(playlist Domain) (Domain, error)
	Get(playlist Domain) ([]Domain, error)
	Delete(playlist Domain) (Domain, error)
}
