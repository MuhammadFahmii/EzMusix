package response

import "EzMusix/bussiness/users"

type Users struct {
	Username string `json:"username" form:"username"`
	Token    string `json:"token"`
}

type UsersRegister struct {
	Username string `json:"username"`
}

type UsersPlaylist struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func FromUsersRegister(usersDomain users.Domain) UsersRegister {
	return UsersRegister{
		Username: usersDomain.Username,
	}
}

func FromUsersLogin(usersDomain users.Domain) Users {
	return Users{
		Username: usersDomain.Username,
		Token:    usersDomain.Token,
	}
}

func ToUsersPlaylist(usersDomain users.Domain) UsersPlaylist {
	return UsersPlaylist{
		Id:       usersDomain.Id,
		Username: usersDomain.Username,
	}
}
