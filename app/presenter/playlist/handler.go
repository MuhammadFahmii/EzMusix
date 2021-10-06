package playlist

import (
	responseHandler "EzMusix/app/presenter"
	"EzMusix/app/presenter/playlist/request"
	"EzMusix/app/presenter/playlist/response"
	"EzMusix/bussiness/playlists"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	playlistUC playlists.Usecase
}

func NewHandler(pl playlists.Usecase) *Presenter {
	return &Presenter{
		playlistUC: pl,
	}
}

func (presenter *Presenter) Insert(c echo.Context) error {
	reqPlaylist := request.Playlist{}
	c.Bind(reqPlaylist)
	domain := request.ToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Insert(domain)
	if err != nil {
		if err.Error() == "please fill all fields" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response.ToAddPlaylist(res))
}

func (presenter *Presenter) Get(c echo.Context) error {
	reqPlaylist := request.Playlist{}
	domain := request.ToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Get(domain)
	if err != nil {
		if err.Error() == "not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resFromDomain := []response.Playlist{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, resFromDomain)
}

func (presenter *Presenter) Delete(c echo.Context) error {
	reqPlaylist := request.DeletePlaylist{}
	c.Bind(reqPlaylist)
	domain := request.DeleteToDomain(reqPlaylist)
	res, err := presenter.playlistUC.Delete(domain)
	if err != nil {
		if err.Error() == "not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, response.FromDomain(res))
}
