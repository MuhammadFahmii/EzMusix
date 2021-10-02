package routes

import (
	"EzMusix/app/presenter/playlist"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	PlaylistHandler playlist.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	p := e.Group("/playlists")
	p.GET("", handler.PlaylistHandler.Get)
}
