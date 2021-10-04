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

func (u *User) toDomain() users.Domain {
	return users.Domain{
		Id:       u.Id,
		Username: u.Username,
		// Playlists: toPlaylistDomain(u.Playlists),
	}
}

// func toPlaylistDomain(u []playlistRepo.Playlist) []playlists.Domain {
// 	playlist := users.Domain{}.Playlists
// 	for _, val := range u {
// 		playlist = append(playlist, playlists.Domain(val))
// 	}
// 	return playlist
// }
