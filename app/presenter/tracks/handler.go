package tracks

import (
	"EzMusix/app/presenter/tracks/request"
	"EzMusix/app/presenter/tracks/response"
	"EzMusix/bussiness/tracks"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	trackUC tracks.Usecase
}

func NewHandler(track tracks.Usecase) *Presenter {
	return &Presenter{
		trackUC: track,
	}
}

func (presenter *Presenter) Get(c echo.Context) error {
	reqTrack := request.Tracks{}
	c.Bind(&reqTrack)
	res, err := presenter.trackUC.Get(reqTrack.TrackName, reqTrack.ArtistName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (presenter *Presenter) AddDetailPlaylist(c echo.Context) error {
	addPlaylist := request.DetailPlaylist{}
	if err := c.Bind(&addPlaylist); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	newDetailTrack := request.ToDetailPlaylist(addPlaylist)
	res, err := presenter.trackUC.AddDetailPlaylist(newDetailTrack)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (presenter *Presenter) DeleteDetailPlaylist(c echo.Context) error {
	addPlaylist := request.DeleteDetailPlaylist{}
	if err := c.Bind(&addPlaylist); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	res, err := presenter.trackUC.DeleteDetailPlaylist(addPlaylist.PlaylistId, addPlaylist.TrackId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response.FromDomain(res))
}
