package users

import (
	"EzMusix/bussiness/users"
	"EzMusix/repository/mysql/playlist"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Playlists []playlist.Playlist
}

func (user *User) toDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		Playlists: user.Playlists,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		Username:  domain.Username,
		Password:  domain.Password,
		Playlists: domain.Playlists,
	}
}
