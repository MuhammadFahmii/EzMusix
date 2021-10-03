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

func (user *User) toDomain() users.Users {
	return users.Users{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		Playlists: user.Playlists,
	}
}

func FromDomain(users users.Users) User {
	return User{
		Username: users.Username,
		Password: users.Password,
	}
}
