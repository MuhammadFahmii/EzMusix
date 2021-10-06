package playlist

import (
	"EzMusix/app/presenter/playlist/request"
	"EzMusix/app/presenter/playlist/response"
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

func (presenter *Presenter) Insert(c echo.Context) error {
	reqPlaylist := request.Playlist{}
	if err := c.Bind(&reqPlaylist); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	domain := request.ToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Insert(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response.FromDomain(res))
}

func (presenter *Presenter) Get(c echo.Context) error {
	reqPlaylist := request.Playlist{}
	domain := request.ToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Get(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	resFromDomain := []response.Playlist{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, *response.FromDomain(val))
	}
	return c.JSON(http.StatusOK, resFromDomain)
}

func (presenter *Presenter) Delete(c echo.Context) error {
	reqPlaylist := request.DeletePlaylist{}
	if err := c.Bind(&reqPlaylist); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	domain := request.DeleteToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Delete(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response.FromDomain(res))
}
