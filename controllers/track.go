package controllers

import (
	"EzMusix/lib/databases"
	"EzMusix/lib/third_party"
	"EzMusix/models/tracks"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTrack(c echo.Context) error {
	tracks := []tracks.Track{}
	if _, err := databases.GetTrack(&tracks); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": tracks,
	})

}

func GetTrackAPI(c echo.Context) error {
	track := tracks.Track{}
	c.Bind(&track)
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": third_party.Search(track.Name, track.ArtistName),
	})
}

func DeleteTrack(c echo.Context) error {
	track := tracks.Track{}
	c.Bind(&track)
	if _, err := databases.DeleteTrack(&track); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": track,
	})
}
