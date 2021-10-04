package users

import (
	"EzMusix/bussiness/users"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DBConn *gorm.DB
}

func NewUserRepo(db *gorm.DB) users.Repository {
	return &UsersRepo{
		DBConn: db,
	}
}

func (usersRepo *UsersRepo) Register(usersDomain users.Domain) (users.Domain, error) {
	usersFromDomain := FromDomain(usersDomain)
	if err := usersRepo.DBConn.Create(&usersFromDomain).Error; err != nil {
		return users.Domain{}, err
	}
	return usersFromDomain.toDomain(), nil
}

func (usersRepo *UsersRepo) Login(usersDomain users.Domain) (users.Domain, error) {
	usersFromDomain := FromDomain(usersDomain)
	if err := usersRepo.DBConn.Where("username=? AND password=?", usersDomain.Username, usersDomain.Password).First(&usersFromDomain).Error; err != nil {
		return users.Domain{}, err
	}
	return usersFromDomain.toDomain(), nil
}
