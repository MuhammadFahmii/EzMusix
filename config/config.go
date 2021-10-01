package config

import (
	"EzMusix/models/playlists"
	"EzMusix/models/tracks"
	"EzMusix/models/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "heinz",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "ez_musix",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&users.User{}, &tracks.Track{}, &playlists.Playlist{})
}
