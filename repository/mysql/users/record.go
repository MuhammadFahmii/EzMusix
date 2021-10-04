package users

import (
	"EzMusix/bussiness/users"
	playlistRepo "EzMusix/repository/mysql/playlist"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Playlists []playlistRepo.Playlist
}

func FromDomain(users users.Domain) User {
	return User{
		Username: users.Username,
		Password: users.Password,
	}
}

func (pl *User) toDomain() users.Domain {
	return users.Domain{
		Id:        pl.Id,
		Username:  pl.Username,
		Password:  pl.Password,
		Playlists: pl.Playlists,
	}
}
