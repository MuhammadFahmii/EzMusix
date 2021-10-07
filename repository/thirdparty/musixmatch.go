package thirdparty

import (
	"EzMusix/bussiness/tracks"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type ThirdParty struct {
	DBConn *gorm.DB
}

func NewThirdParty(conn *gorm.DB) tracks.ThirdParty {
	return &ThirdParty{
		DBConn: conn,
	}
}

func (tracksRepo *ThirdParty) Get(trackName, artistName string) (tracks.Domain, error) {
	const API_KEY = "4b69788a296733b380069f770a174a89"
	trackName = strings.Replace(trackName, " ", "-", -1)
	artistName = strings.Replace(artistName, " ", "-", -1)
	var url = fmt.Sprintf(`http://api.musixmatch.com/ws/1.1/track.search?apikey=%s&page_size=1&q_track=%s&q_artist=%s`, API_KEY, trackName, artistName)
	res, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	response := Response{}
	if err := json.Unmarshal(responseData, &response); err != nil {
		return tracks.Domain{}, err
	}
	if len(response.Message.Body.TrackList) == 0 {
		return tracks.Domain{}, errors.New("not found")
	}
	return response.toDomain(), nil
}
