package tracks

import (
	"EzMusix/app/presenter/tracks/request"
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
	domain := reqTrack.ToDomain()
	res, err := presenter.trackUC.Get(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
