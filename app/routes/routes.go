package routes

import (
	"EzMusix/app/middlewares"
	"EzMusix/app/presenter/playlist"
	"EzMusix/app/presenter/tracks"
	"EzMusix/app/presenter/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	DBLog           middlewares.LogMiddleware
	JWTMiddleware   middleware.JWTConfig
	PlaylistHandler playlist.Presenter
	TrackHandler    tracks.Presenter
	UsersHandler    users.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(handler.DBLog.Log)
	u := e.Group("/users")
	u.POST("/register", handler.UsersHandler.Register)
	u.POST("/login", handler.UsersHandler.Login)
	u.GET("", handler.UsersHandler.GetAllUsers)

	p := e.Group("/playlists")
	p.GET("", handler.PlaylistHandler.Get, middleware.JWTWithConfig(handler.JWTMiddleware))
	p.POST("", handler.PlaylistHandler.Insert)
	p.DELETE("/:id", handler.PlaylistHandler.Delete)

	t := e.Group("/tracks")
	t.GET("", handler.TrackHandler.Get)

	dp := e.Group("/detailPlaylist")
	dp.POST("", handler.TrackHandler.AddDetailPlaylist)
	dp.DELETE("/:playlist_id/:track_id", handler.TrackHandler.DeleteDetailPlaylist)
}
