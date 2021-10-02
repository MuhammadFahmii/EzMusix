package thirdparty

import (
	"EzMusix/bussiness/tracks"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type TracksRepo struct {
	DBConn *gorm.DB
}

func NewTracksRepo(conn *gorm.DB) tracks.Repository {
	return &TracksRepo{
		DBConn: conn,
	}
}

func (tracksRepo *TracksRepo) Get(track tracks.Domain) (tracks.Domain, error) {
	const API_KEY = "4b69788a296733b380069f770a174a89"
	trackName := strings.Replace(track.Name, " ", "-", -1)
	artistName := strings.Replace(track.ArtistName, " ", "-", -1)
	var url = fmt.Sprintf(`http://api.musixmatch.com/ws/1.1/track.search?apikey=%s&page_size=1&q_track=%s&q_artist=%s`, API_KEY, trackName, artistName)
	res, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	response := Response{}
	json.Unmarshal(responseData, &response)
	return response.toDomain(), nil
}
