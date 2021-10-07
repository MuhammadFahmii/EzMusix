package mysql

import (
	"EzMusix/repository/mysql/comments"
	"EzMusix/repository/mysql/playlist"
	"EzMusix/repository/mysql/tracks"
	"EzMusix/repository/mysql/users"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	config := map[string]string{
		"DB_Username": viper.GetString("database.username"),
		"DB_Password": viper.GetString("database.password"),
		"DB_Port":     viper.GetString("database.port"),
		"DB_Host":     viper.GetString("database.host"),
		"DB_Name":     viper.GetString("database.name"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&users.User{},
		&playlist.Playlist{},
		&tracks.Track{},
		&comments.Comments{},
	)

	return DB
}
