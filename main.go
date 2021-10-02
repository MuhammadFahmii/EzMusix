package main

import (
	"EzMusix/app/middlewares"
	playlistHandler "EzMusix/app/presenter/playlist"
	tracksHandler "EzMusix/app/presenter/tracks"
	usersHandler "EzMusix/app/presenter/users"
	"EzMusix/app/routes"
	playlistUsecase "EzMusix/bussiness/playlist"
	tracksUsecase "EzMusix/bussiness/tracks"
	usersUsecase "EzMusix/bussiness/users"
	playlistRepo "EzMusix/repository/mysql/playlist"
	usersRepo "EzMusix/repository/mysql/users"
	trackRepo "EzMusix/repository/thirdparty"
	"fmt"

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
		&usersRepo.User{},
	)

	return DB
}

func main() {
	e := echo.New()
	db := InitDB()
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       "12345",
		ExpiresDuration: 1,
	}
	// Users
	usersRepo := usersRepo.NewUserRepo(db)
	usersUsecase := usersUsecase.NewUserUsecase(usersRepo, &configJWT)
	usersHandler := usersHandler.NewHandler(usersUsecase)

	// Playlists
	playlistRepo := playlistRepo.NewPlaylistRepo(db)
	playlistUsecase := playlistUsecase.NewPlaylistUsecase(playlistRepo)
	playlistHandler := playlistHandler.NewHandler(playlistUsecase)

	// Tracks
	tracksRepo := trackRepo.NewTracksRepo(db)
	tracksUsecase := tracksUsecase.NewTracksUsecase(tracksRepo)
	tracksHandler := tracksHandler.NewHandler(tracksUsecase)
	routesInit := routes.HandlerList{
		JWTMiddleware:   configJWT.Init(),
		PlaylistHandler: *playlistHandler,
		TrackHandler:    *tracksHandler,
		UsersHandler:    *usersHandler,
	}
	routesInit.RouteRegister(e)
	e.Start(":8000")
}
