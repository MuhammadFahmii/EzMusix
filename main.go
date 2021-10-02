package main

import (
	playlistHandler "EzMusix/app/presenter/playlist"
	"EzMusix/app/routes"
	playlistUsecase "EzMusix/bussiness/playlist"
	playlistRepo "EzMusix/repository/mysql/playlist"
	"fmt"

	trackRepo "EzMusix/repository/mysql/tracks"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "heinz",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "ez_musix",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&playlistRepo.Playlist{},
		&trackRepo.Track{},
	)

	return DB
}

func main() {
	e := echo.New()
	db := InitDB()
	playlistRepo := playlistRepo.NewPlaylistRepo(db)
	playlistUsecase := playlistUsecase.NewPlaylistUsecase(playlistRepo)
	playlistHandler := playlistHandler.NewHandler(playlistUsecase)
	routesInit := routes.HandlerList{
		PlaylistHandler: *playlistHandler,
	}
	routesInit.RouteRegister(e)
	e.Start(":8000")
}
