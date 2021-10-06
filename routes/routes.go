package routes

import (
	"EzMusix/constants"
	"EzMusix/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RemoveTrailingSlash())
	jwt := middleware.JWT([]byte(constants.SECRET_JWT))
	u := e.Group("/users")
	u.POST("/register", controllers.RegisterUser)
	u.POST("/login", controllers.LoginUser)
	u.GET("", controllers.GetUser, jwt)

	g := e.Group("/tracks")
	g.GET("", controllers.GetTrack)
	g.GET("/3rd_party", controllers.GetTrackAPI)
	g.DELETE("/:trackId", controllers.DeleteTrack)

	p := e.Group("/playlists")
	p.POST("", controllers.AddPlaylist)
	p.GET("", controllers.GetPlaylist)
	p.DELETE("/:playlist_id", controllers.DeletePlaylist)

	tp := e.Group("/detailPlaylist")
	tp.POST("", controllers.AddDetailPlaylist)
	tp.DELETE("/:playlist_id/:track_id", controllers.DeleteDetailPlaylist)

	return e
}
