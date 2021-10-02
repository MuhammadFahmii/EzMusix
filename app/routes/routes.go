package routes

import (
	"EzMusix/app/presenter/playlist"
	"EzMusix/app/presenter/tracks"
	"EzMusix/app/presenter/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware   middleware.JWTConfig
	PlaylistHandler playlist.Presenter
	TrackHandler    tracks.Presenter
	UsersHandler    users.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	u := e.Group("/users")
	u.POST("/register", handler.UsersHandler.Register)
	u.POST("/login", handler.UsersHandler.Login)

	p := e.Group("/playlists")
	e.Use(middleware.RemoveTrailingSlash())
	p.GET("", handler.PlaylistHandler.Get, middleware.JWTWithConfig(handler.JWTMiddleware))
	p.POST("", handler.PlaylistHandler.Insert)
	p.DELETE("/:id", handler.PlaylistHandler.Delete)

	t := e.Group("/tracks")
	t.GET("", handler.TrackHandler.Get)
}
