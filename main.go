package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
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
	Id         int    `gorm:"primarykey" json:"track_id" form:"track_id"`
	Name       string `json:"track_name" form:"track_name"`
	AlbumName  string `json:"album_name" form:"album_name"`
	ArtistName string `json:"artist_name" form:"artist_name"`
}

type Playlist struct {
	Id     int     `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name" form:"name"`
	UserID int     `json:"user_id" form:"user_id"`
	Tracks []Track `gorm:"many2many:detail_playlist constraint:OnDelete:CASCADE"`
}

type DetailPlaylist struct {
	PlaylistId int    `json:"playlist_id" form:"playlist_id"`
	TrackName  string `json:"track_name" form:"track_name"`
	ArtistName string `json:"artist_name" form:"artist_name"`
}

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Playlists []Playlist
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
	DB.AutoMigrate(&User{}, &Track{}, &Playlist{})
}

func main() {
	InitDB()
	e := echo.New()

	u := e.Group("/users")
	u.POST("/register", RegisterUser)
	u.POST("/login", LoginUser)
	u.GET("", GetUser)

	g := e.Group("/tracks")
	g.GET("", GetTrack)
	g.DELETE("/:trackId", DeleteTrack)

	p := e.Group("/playlists")
	p.POST("", AddPlaylist)
	p.GET("", GetPlaylist)
	p.DELETE("/:playlist_id", DeletePlaylist)

	tp := e.Group("/detailPlaylist")
	tp.POST("", AddDetailPlaylist)
	tp.DELETE("/:playlist_id/:track_id", DeleteDetailPlaylist)

	e.Logger.Fatal(e.Start(":8000"))
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

/*
User Method
*/
func RegisterUser(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg": "Success",
	})
}
func LoginUser(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := DB.Debug().Where("username=? AND password=?", user.Username, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{})
}
func GetUser(c echo.Context) error {
	user := []User{}
	if err := DB.Debug().Preload("Playlists").Find(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": user,
	})
}

/*
Track Method
*/
func GetTrack(c echo.Context) error {
	track := c.QueryParam("track_name")
	artist := c.QueryParam("artist_name")
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

/*
Playlist Method
*/
func AddPlaylist(c echo.Context) error {
	playlist := Playlist{}
	c.Bind(&playlist)
	if err := DB.Debug().Create(&playlist).Error; err != nil {
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

func DeletePlaylist(c echo.Context) error {
	playlistId, _ := strconv.Atoi(c.Param("playlist_id"))
	DB.Where("id=?", playlistId).Delete(&Playlist{})
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "success",
	})
}

/*
DetailPlaylist Method
*/
func AddDetailPlaylist(c echo.Context) error {
	detailPlaylist := DetailPlaylist{}
	c.Bind(&detailPlaylist)
	var newTrack Track
	newTrack = Search(detailPlaylist.TrackName, detailPlaylist.ArtistName)
	DB.Model(&Playlist{Id: detailPlaylist.PlaylistId}).Association("Tracks").Append(&newTrack)
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":  "Success",
		"data": newTrack,
	})
}

func DeleteDetailPlaylist(c echo.Context) error {
	playlistId, _ := strconv.Atoi(c.Param("playlist_id"))
	trackId, _ := strconv.Atoi(c.Param("track_id"))
	DB.Model(&Playlist{Id: playlistId}).Association("Tracks").Delete(&Track{Id: trackId})
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "success",
	})
}
