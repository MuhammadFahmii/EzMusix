package databases

import (
	"EzMusix/config"
	"EzMusix/models/users"
)

func LoginUser(user *users.User) (interface{}, error) {
	if err := config.DB.Where("username=? AND password=?", user.Username, user.Password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func AddUser(user *users.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser(user *[]users.User) (interface{}, error) {
	if err := config.DB.Preload("Playlists").Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
