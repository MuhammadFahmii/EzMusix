package routes

import (
	"EzMusix/app/middlewares"
	"EzMusix/app/presenter/comments"
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
	CommentsHandler comments.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(handler.DBLog.Log)
	u := e.Group("/users")
	u.POST("/register", handler.UsersHandler.Register)
	u.POST("/login", handler.UsersHandler.Login)
	u.GET("", handler.UsersHandler.GetAllUsers, middleware.JWTWithConfig(handler.JWTMiddleware))
	u.PUT("", handler.UsersHandler.UpdateUsers, middleware.JWTWithConfig(handler.JWTMiddleware))

	p := e.Group("/playlists", middleware.JWTWithConfig(handler.JWTMiddleware))
	p.GET("", handler.PlaylistHandler.Get)
	p.POST("", handler.PlaylistHandler.Insert)
	p.DELETE("/:id", handler.PlaylistHandler.Delete)

	c := e.Group("/comments", middleware.JWTWithConfig(handler.JWTMiddleware))
	c.GET("", handler.CommentsHandler.Get)
	c.POST("", handler.CommentsHandler.Insert)
	c.DELETE("/:id", handler.CommentsHandler.Delete)

	t := e.Group("/tracks")
	t.GET("", handler.TrackHandler.Get)

	dp := e.Group("/detailPlaylist", middleware.JWTWithConfig(handler.JWTMiddleware))
	dp.POST("", handler.TrackHandler.AddDetailPlaylist)
	dp.DELETE("/:playlist_id/:track_id", handler.TrackHandler.DeleteDetailPlaylist)
}
