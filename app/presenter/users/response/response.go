package response

import "EzMusix/bussiness/users"

type Users struct {
	Username string `json:"username" form:"username"`
	Token    string `json:"token"`
}

func FromDomain(usersDomain users.Domain) Users {
	return Users{
		Username: usersDomain.Username,
		Token:    usersDomain.Token,
	}
}
