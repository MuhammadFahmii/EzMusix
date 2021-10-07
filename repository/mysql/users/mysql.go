package users

import (
	"EzMusix/bussiness/users"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var mysqlErr *mysql.MySQLError

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
	usersRepo.DBConn.Where("username=?", usersDomain.Username).First(&usersFromDomain)
	if err := usersRepo.DBConn.Create(&usersFromDomain).Error; err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return users.Domain{}, errors.New("duplicate entry")
		}
		return users.Domain{}, err
	}
	return usersFromDomain.toDomain(), nil
}

func (usersRepo *UsersRepo) Login(usersDomain users.Domain) (users.Domain, error) {
	usersFromDomain := FromDomain(usersDomain)
	if err := usersRepo.DBConn.Where("username=? AND password=?", usersDomain.Username, usersDomain.Password).First(&usersFromDomain).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.Domain{}, errors.New("record not found")
		}
		return users.Domain{}, err
	}
	return usersFromDomain.toDomain(), nil
}
func (usersRepo *UsersRepo) GetAllUsers(usersDomain users.Domain) ([]users.Domain, error) {
	rec := []User{}
	if err := usersRepo.DBConn.Debug().Preload("Playlists.Tracks").Find(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []users.Domain{}, errors.New("record not found")
		}
		return []users.Domain{}, err
	}
	var domainPlaylist []users.Domain
	for _, val := range rec {
		domainPlaylist = append(domainPlaylist, val.toDomain())
	}
	return domainPlaylist, nil
}

func (usersRepo *UsersRepo) UpdateUsers(usersDomain users.Domain) (users.Domain, error) {
	rec := FromDomain(usersDomain)
	if err := usersRepo.DBConn.Debug().Save(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.Domain{}, errors.New("record not found")
		}
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}
