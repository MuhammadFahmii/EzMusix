package tracks

import (
	responseHandler "EzMusix/app/presenter"
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
	res, err := presenter.trackUC.Get(reqTrack.TrackName, reqTrack.ArtistName)
	if err != nil {
		if err.Error() == "not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, res)
}

func (presenter *Presenter) AddDetailPlaylist(c echo.Context) error {
	addPlaylist := request.DetailPlaylist{}
	c.Bind(&addPlaylist)
	newDetailTrack := request.ToDetailPlaylist(addPlaylist)
	res, err := presenter.trackUC.AddTrackPlaylist(newDetailTrack)
	if err != nil {
		responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusCreated, res)
}

func (presenter *Presenter) DeleteDetailPlaylist(c echo.Context) error {
	addPlaylist := request.DeleteDetailPlaylist{}
	c.Bind(&addPlaylist)
	res, err := presenter.trackUC.DeleteTrackPlaylist(addPlaylist.PlaylistId, addPlaylist.TrackId)
	if err != nil {
		if err.Error() == "please fill all param" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)

	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, res)
}
