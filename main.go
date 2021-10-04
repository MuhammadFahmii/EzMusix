package main

import (
	"EzMusix/app/middlewares"
	playlistHandler "EzMusix/app/presenter/playlist"
	tracksHandler "EzMusix/app/presenter/tracks"
	usersHandler "EzMusix/app/presenter/users"
	"EzMusix/app/routes"
	playlistUsecase "EzMusix/bussiness/playlists"
	tracksUsecase "EzMusix/bussiness/tracks"
	usersUsecase "EzMusix/bussiness/users"
	"EzMusix/repository/mongodb"
	"EzMusix/repository/mysql"
	playlistRepo "EzMusix/repository/mysql/playlist"
	usersRepo "EzMusix/repository/mysql/users"
	trackRepo "EzMusix/repository/thirdparty"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("app/config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	db := mysql.InitDB()
	dbLog := middlewares.LogMiddleware{
		DBLog: mongodb.InitLog(),
	}
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("secret"),
		ExpiresDuration: viper.GetInt("expired"),
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
		DBLog:           dbLog,
		JWTMiddleware:   configJWT.Init(),
		PlaylistHandler: *playlistHandler,
		TrackHandler:    *tracksHandler,
		UsersHandler:    *usersHandler,
	}
	routesInit.RouteRegister(e)
	e.Start(":8000")
}
