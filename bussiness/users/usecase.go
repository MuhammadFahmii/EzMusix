package users

import (
	"EzMusix/app/middlewares"
	"errors"
)

type UserUsecase struct {
	userRepo Repository
	jwtAuth  *middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		userRepo: repo,
		jwtAuth:  jwtAuth,
	}
}
func (uc *UserUsecase) Register(usersDomain Domain) (Domain, error) {
	if usersDomain.Username == "" {
		return Domain{}, errors.New("Username empty")
	}
	user, err := uc.userRepo.Register(usersDomain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Login(usersDomain Domain) (Domain, error) {
	if usersDomain.Username == "" || usersDomain.Password == "" {
		return Domain{}, errors.New("Username or password is empty")
	}
	user, err := uc.userRepo.Login(usersDomain)
	if err != nil {
		return Domain{}, err
	}
	user.Token, err = uc.jwtAuth.GenerateToken(user.Id)
	return user, nil
}
