package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Response struct {
	Message Message `json:"message"`
}
type Message struct {
	Body Body `json:"body"`
}
type Body struct {
	TrackList []TrackList `json:"track_list"`
}
type TrackList struct {
	Track Track `json:"track"`
}
type Track struct {
	Id         int    `gorm:"primarykey" json:"track_id"`
	Name       string `json:"track_name"`
	AlbumName  string `json:"album_name"`
	ArtistName string `json:"artist_name"`
}

type Playlist struct {
	Id     int     `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name"`
	Tracks []Track `gorm:"many2many:detail_playlist"`
}

func InitDB() {
	config := map[string]string{
		"DB_Username": "heinz",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "ez_musix",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&Track{}, &Playlist{})
}

func main() {
	InitDB()
	e := echo.New()
	g := e.Group("/tracks")
	g.GET("", GetTrack)
	g.DELETE("/:trackId", DeleteTrack)

	p := e.Group("/playlists")
	p.POST("", AddPlaylist)
	p.GET("", GetPlaylist)

	tp := e.Group("/detailPlaylist")
	tp.POST("", AddDetailPlaylist)

	e.Logger.Fatal(e.Start(":8000"))
}

func AddDetailPlaylist(c echo.Context) error {
	playlist := Playlist{}
	playlistId, _ := strconv.Atoi(c.FormValue("playlist_id"))
	track := c.FormValue("q_track")
	artist := c.FormValue("q_artist")
	var newTrack Track
	newTrack = Search(track, artist)
	DB.Where("id=?", playlistId).Find(&playlist).Association("Tracks").Append(&newTrack)
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":  "Success",
		"data": newTrack,
	})
}

func Search(trackName, artistName string) Track {
	const API_KEY = "4b69788a296733b380069f770a174a89"
	trackName = strings.Replace(trackName, " ", "-", -1)
	artistName = strings.Replace(artistName, " ", "-", -1)
	var url = fmt.Sprintf(`http://api.musixmatch.com/ws/1.1/track.search?apikey=%s&page_size=1&q_track=%s&q_artist=%s`, API_KEY, trackName, artistName)
	res, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Message.Body.TrackList[0].Track
}

func GetTrack(c echo.Context) error {
	track := c.QueryParam("q_track")
	artist := c.QueryParam("q_artist")
	if track == "" && artist == "" {
		track := []Track{}
		if err := DB.Find(&track).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"msg": "failed",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"msg":  "Success",
			"data": track,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": Search(track, artist),
	})
}

func DeleteTrack(c echo.Context) error {
	trackId, _ := strconv.Atoi(c.Param("trackId"))
	track := Track{}
	if err := DB.Where("track_id = ?", trackId).Delete(&track).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "Success",
	})
}

func AddPlaylist(c echo.Context) error {
	playlist := Playlist{}
	c.Bind(&playlist)
	if err := DB.Create(&playlist).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":  "Success",
		"data": playlist,
	})
}

func GetPlaylist(c echo.Context) error {
	playlist := Playlist{}
	if err := DB.Preload("Tracks").Find(&playlist).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": playlist,
	})
}
