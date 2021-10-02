package users

// import (
// 	"EzMusix/bussiness/users"
// 	"EzMusix/models/playlists"
// )

type User struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Password string
}

// func (user *User) ToDomain() users.Domain {
// 	return users.Domain{
// 		Id:        user.Id,
// 		Username:  user.Username,
// 		Password:  user.Password,
// 		Playlists: user.Playlists,
// 	}
// }

// func fromDomain(domain users.Domain) User {
// 	return User{
// 		Id:        domain.Id,
// 		Username:  domain.Username,
// 		Password:  domain.Password,
// 		Playlists: domain.Playlists,
// 	}
// }
