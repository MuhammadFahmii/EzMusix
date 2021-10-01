package controllers

import (
	"EzMusix/lib/databases"
	"EzMusix/models/playlists"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddPlaylist(c echo.Context) error {
	playlist := playlists.Playlist{}
	c.Bind(&playlist)
	if _, err := databases.AddPlaylist(&playlist); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":  "Success",
		"data": playlist,
	})
}

func GetPlaylist(c echo.Context) error {
	playlist := []playlists.Playlist{}
	if _, err := databases.GetPlaylist(&playlist); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": playlist,
	})
}

func DeletePlaylist(c echo.Context) error {
	playlist := playlists.Playlist{}
	c.Bind(&playlist)
	if _, err := databases.DeletePlaylist(&playlist); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "success",
		"data": playlist,
	})
}

/*
DetailPlaylist Method
*/
func AddDetailPlaylist(c echo.Context) error {
	detailPlaylist := playlists.DetailPlaylist{}
	c.Bind(&detailPlaylist)
	if _, err := databases.AddDetailPlaylist(&detailPlaylist); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "success",
		"data": detailPlaylist,
	})
}

func DeleteDetailPlaylist(c echo.Context) error {
	playlistId, _ := strconv.Atoi(c.Param("playlist_id"))
	trackId, _ := strconv.Atoi(c.Param("track_id"))
	res, err := databases.DeleteDetailPlaylist(playlistId, trackId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "success",
		"data": res,
	})
}
