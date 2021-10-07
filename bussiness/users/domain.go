package users

import (
	"EzMusix/repository/mysql/playlist"
	"time"
)

type Domain struct {
	Id        int
	Username  string
	Password  string
	Token     string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Playlists []playlist.Playlist
}

type Usecase interface {
	Register(Domain) (Domain, error)
	Login(Domain) (Domain, error)
	GetAllUsers(Domain) ([]Domain, error)
	UpdateUsers(Domain) (Domain, error)
}

type Repository interface {
	Register(Domain) (Domain, error)
	Login(Domain) (Domain, error)
	GetAllUsers(Domain) ([]Domain, error)
	UpdateUsers(Domain) (Domain, error)
}
