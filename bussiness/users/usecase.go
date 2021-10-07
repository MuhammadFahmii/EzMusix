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
	if usersDomain.Username == "" || usersDomain.Password == "" {
		return Domain{}, errors.New("please fill all fields")
	}
	usersDomain.Password = helpers.Hash(usersDomain.Password)
	user, err := uc.userRepo.Register(usersDomain)
	if err != nil {
		if err.Error() == "duplicate entry" {
			return Domain{}, errors.New("duplicate entry")
		}
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Login(usersDomain Domain) (Domain, error) {
	if usersDomain.Username == "" || usersDomain.Password == "" {
		return Domain{}, errors.New("please fill all fields")
	}
	usersDomain.Password = helpers.Hash(usersDomain.Password)
	user, err := uc.userRepo.Login(usersDomain)
	if err != nil {
		if err.Error() == "record not found" {
			return Domain{}, errors.New("record not found")
		}
		return Domain{}, err
	}
	user.Token, _ = uc.jwtAuth.GenerateToken(user.Id, user.Status)
	return user, nil
}

func (uc *UserUsecase) GetAllUsers(usersDomain Domain) ([]Domain, error) {
	res, err := uc.userRepo.GetAllUsers(usersDomain)
	if err != nil {
		if err.Error() == "record not found" {
			return []Domain{}, errors.New("record not found")
		}
		return []Domain{}, err
	}
	return res, nil
}

func (uc *UserUsecase) UpdateUsers(usersDomain Domain) (Domain, error) {
	res, err := uc.userRepo.UpdateUsers(usersDomain)
	if err != nil {
		if err.Error() == "record not found" {
			return Domain{}, errors.New("record not found")
		}
		return Domain{}, err
	}
	return res, nil
}
