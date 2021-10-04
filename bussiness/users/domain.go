package users

import (
	"EzMusix/repository/mysql/playlist"
)

type Domain struct {
	Id        int
	Username  string
	Password  string
	Token     string
	Status    int
	Playlists []playlist.Playlist
}

type Usecase interface {
	Register(Domain) (Domain, error)
	Login(Domain) (Domain, error)
	GetAllUsers(Domain) ([]Domain, error)
}

type Repository interface {
	Register(Domain) (Domain, error)
	Login(Domain) (Domain, error)
	GetAllUsers(Domain) ([]Domain, error)
}
