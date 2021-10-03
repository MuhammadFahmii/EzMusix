package users

import (
	"errors"
)

type UserUsecase struct {
	userRepo Repository
	// jwtAuth  *middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository) Usecase {
	return &UserUsecase{
		userRepo: repo,
		// jwtAuth:  jwtAuth,
	}
}
func (uc *UserUsecase) Register(usersDomain Users) (Users, error) {
	if usersDomain.Username == "" {
		return Users{}, errors.New("username empty")
	}
	if usersDomain.Password == "" {
		return Users{}, errors.New("password empty")
	}
	user, err := uc.userRepo.Register(usersDomain)
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Login(usersDomain Users) (Users, error) {
	if usersDomain.Username == "" {
		return Users{}, errors.New("password is empty")
	}
	if usersDomain.Password == "" {
		return Users{}, errors.New("password is empty")
	}
	user, err := uc.userRepo.Login(usersDomain)
	if err != nil {
		return Users{}, err
	}
	// user.Token, err = uc.jwtAuth.GenerateToken(user.Id)
	return user, nil
}
