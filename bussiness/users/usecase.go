package users

import (
	"EzMusix/app/middlewares"
	"EzMusix/helpers"
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
		return Domain{}, errors.New("username empty")
	}
	if usersDomain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	usersDomain.Password = helpers.Hash(usersDomain.Password)
	user, err := uc.userRepo.Register(usersDomain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Login(usersDomain Domain) (Domain, error) {
	if usersDomain.Username == "" {
		return Domain{}, errors.New("password is empty")
	}
	if usersDomain.Password == "" {
		return Domain{}, errors.New("password is empty")
	}
	usersDomain.Password = helpers.Hash(usersDomain.Password)
	user, err := uc.userRepo.Login(usersDomain)
	if err != nil {
		return Domain{}, err
	}
	user.Token, _ = uc.jwtAuth.GenerateToken(user.Id, user.Status)
	return user, nil
}

func (uc *UserUsecase) GetAllUsers(usersDomain Domain) ([]Domain, error) {
	res, err := uc.userRepo.GetAllUsers(usersDomain)
	if err != nil {
		return []Domain{}, nil
	}
	return res, nil
}
