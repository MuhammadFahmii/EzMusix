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
func (usersRepo *UsersRepo) GetAllUsers(usersDomain users.Domain) ([]users.Domain, error) {
	rec := []User{}
	if err := usersRepo.DBConn.Debug().Preload("Playlists").Find(&rec).Error; err != nil {
		return []users.Domain{}, err
	}
	var domainPlaylist []users.Domain
	for _, val := range rec {
		domainPlaylist = append(domainPlaylist, val.toDomain())
	}
	return domainPlaylist, nil
}
