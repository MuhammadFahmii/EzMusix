package users

import (
	"EzMusix/bussiness/users"
	"EzMusix/repository/mysql/comments"
	playlistRepo "EzMusix/repository/mysql/playlist"
	"time"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Username  string
	Password  string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Playlists []playlistRepo.Playlist
	Comments  []comments.Comments
}

func FromDomain(users users.Domain) User {
	return User{
		Id:        users.Id,
		Username:  users.Username,
		Password:  users.Password,
		Status:    users.Status,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
	}
}

func (pl *User) toDomain() users.Domain {
	return users.Domain{
		Id:        pl.Id,
		Username:  pl.Username,
		Password:  pl.Password,
		Status:    pl.Status,
		CreatedAt: pl.CreatedAt,
		UpdatedAt: pl.UpdatedAt,
	}
}
