package response

import "EzMusix/bussiness/users"

type Users struct {
	Username string `json:"username" form:"username"`
	Token    string `json:"token"`
}

type UsersRegister struct {
	Username string `json:"username"`
}

func FromUsersRegister(usersDomain users.Users) UsersRegister {
	return UsersRegister{
		Username: usersDomain.Username,
	}
}

func FromUsersLogin(usersDomain users.Users) Users {
	return Users{
		Username: usersDomain.Username,
		Token:    usersDomain.Token,
	}
}
