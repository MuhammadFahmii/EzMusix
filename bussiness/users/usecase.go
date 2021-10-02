package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	userRepo       Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeOut time.Duration) Usecase {
	return &UserUsecase{
		userRepo:       repo,
		contextTimeout: timeOut,
	}
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	user, err := uc.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
