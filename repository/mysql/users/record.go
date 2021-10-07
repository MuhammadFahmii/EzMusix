package users

import (
	"EzMusix/bussiness/users"
	playlistRepo "EzMusix/repository/mysql/playlist"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Status    int
	Playlists []playlistRepo.Playlist
}

func FromDomain(users users.Domain) User {
	return User{
		Username: users.Username,
		Password: users.Password,
		Status:   users.Status,
	}
}

func (pl *User) toDomain() users.Domain {
	return users.Domain{
		Id:       pl.Id,
		Username: pl.Username,
		Password: pl.Password,
		Status:   pl.Status,
	}
}
