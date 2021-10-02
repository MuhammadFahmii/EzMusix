package playlist

import (
	"EzMusix/app/presenter/playlist/request"
	"EzMusix/bussiness/playlist"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	playlistUC playlist.Usecase
}

func NewHandler(pl playlist.Usecase) *Presenter {
	return &Presenter{
		playlistUC: pl,
	}
}

func (presenter *Presenter) Get(c echo.Context) error {
	reqPlaylist := request.Playlist{}
	if err := c.Bind(&reqPlaylist); err != nil {
		c.JSON(http.StatusBadRequest, "Something Wrong")
	}
	domain := request.ToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Get(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
