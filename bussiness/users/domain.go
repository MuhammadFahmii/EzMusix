package users

import (
	"context"
)

type Domain struct {
	Id           int
	Username     string
	Password     string
	Token        string
	PlaylistId   int
	PlaylistName string
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email, password string) (Domain, error)
}
