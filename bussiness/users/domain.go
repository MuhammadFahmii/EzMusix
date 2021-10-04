package users

import "EzMusix/bussiness/playlists"

type Domain struct {
	Id        int
	Username  string
	Password  string
	Token     string
	Playlists []playlists.Domain
}

type Usecase interface {
	Register(Domain) (Domain, error)
	Login(Domain) (Domain, error)
}

type Repository interface {
	Register(Domain) (Domain, error)
	Login(Domain) (Domain, error)
}
