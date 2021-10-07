package response

import (
	"EzMusix/bussiness/users"
	"time"
)

type Users struct {
	Username  string    `json:"username" form:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BasicUsersResponse struct {
	Username string `json:"username"`
}

type UsersPlaylist struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromUsersUpdate(usersDomain users.Domain) BasicUsersResponse {
	return BasicUsersResponse{
		Username: usersDomain.Username,
	}
}

func FromUsersRegister(usersDomain users.Domain) BasicUsersResponse {
	return BasicUsersResponse{
		Username: usersDomain.Username,
	}
}

func FromUsersLogin(usersDomain users.Domain) Users {
	return Users{
		Username:  usersDomain.Username,
		Token:     usersDomain.Token,
		CreatedAt: usersDomain.CreatedAt,
		UpdatedAt: usersDomain.UpdatedAt,
	}
}

func ToUsersPlaylist(usersDomain users.Domain) UsersPlaylist {
	return UsersPlaylist{
		Id:        usersDomain.Id,
		Username:  usersDomain.Username,
		CreatedAt: usersDomain.CreatedAt,
		UpdatedAt: usersDomain.CreatedAt,
	}
}
