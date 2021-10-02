package request

import (
	"EzMusix/bussiness/users"
)

type Users struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (usersReq Users) ToDomain() users.Domain {
	return users.Domain{
		Username: usersReq.Username,
		Password: usersReq.Password,
	}
}
