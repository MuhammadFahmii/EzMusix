package main

import (
	"EzMusix/app/middlewares"
	commentsHandler "EzMusix/app/presenter/comments"
	playlistHandler "EzMusix/app/presenter/playlist"
	tracksHandler "EzMusix/app/presenter/tracks"
	usersHandler "EzMusix/app/presenter/users"
	"EzMusix/app/routes"
	commentsUsecase "EzMusix/bussiness/comments"
	playlistUsecase "EzMusix/bussiness/playlists"
	tracksUsecase "EzMusix/bussiness/tracks"
	usersUsecase "EzMusix/bussiness/users"

	"EzMusix/repository/mongodb"
	"EzMusix/repository/mysql"
	commentsRepo "EzMusix/repository/mysql/comments"
	playlistRepo "EzMusix/repository/mysql/playlist"
	trackRepo "EzMusix/repository/mysql/tracks"
	usersRepo "EzMusix/repository/mysql/users"
	"EzMusix/repository/thirdparty"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
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
	mongoDB := mongodb.InitLog()
	dbLog := middlewares.LogMiddleware{
		DBLog: mongoDB,
	}
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("jwt.secret"),
		ExpiresDuration: viper.GetInt("jwt.expired"),
	}
	// Users
	usersRepo := usersRepo.NewUserRepo(db)
	usersUsecase := usersUsecase.NewUserUsecase(usersRepo, &configJWT)
	usersHandler := usersHandler.NewHandler(usersUsecase)

	// Playlists
	playlistRepo := playlistRepo.NewPlaylistRepo(db)
	playlistUsecase := playlistUsecase.NewPlaylistUsecase(playlistRepo)
	playlistHandler := playlistHandler.NewHandler(playlistUsecase)

	commentsRepo := commentsRepo.NewCommentsRepo(db)
	commentsUsecase := commentsUsecase.NewCommentsUsecase(commentsRepo)
	commentsHandler := commentsHandler.NewHandler(commentsUsecase)

	// Tracks
	tracksRepo := trackRepo.NewTracksRepo(db)
	thirdParty := thirdparty.NewThirdParty(db)
	tracksUsecase := tracksUsecase.NewTracksUsecase(tracksRepo, thirdParty)
	tracksHandler := tracksHandler.NewHandler(tracksUsecase)
	routesInit := routes.HandlerList{
		DBLog:           dbLog,
		JWTMiddleware:   configJWT.Init(),
		PlaylistHandler: *playlistHandler,
		TrackHandler:    *tracksHandler,
		UsersHandler:    *usersHandler,
		CommentsHandler: *commentsHandler,
	}
	routesInit.RouteRegister(e)
	e.Start(viper.GetString("server.address"))
}
