package users

import (
	"EzMusix/repository/mysql/playlist"
)

type Users struct {
	Id        int
	Username  string
	Password  string
	Token     string
	Playlists []playlist.Playlist
}

type Usecase interface {
	Register(Users) (Users, error)
	Login(Users) (Users, error)
}

type Repository interface {
	Register(Users) (Users, error)
	Login(Users) (Users, error)
}
