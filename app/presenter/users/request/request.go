package request

import (
	"EzMusix/bussiness/users"
)

type Users struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (usersReq Users) ToDomain() users.Users {
	return users.Users{
		Username: usersReq.Username,
		Password: usersReq.Password,
	}
}
