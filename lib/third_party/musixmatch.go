package third_party

import (
	"EzMusix/models/tracks"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Search(trackName, artistName string) tracks.Track {
	const API_KEY = "4b69788a296733b380069f770a174a89"
	trackName = strings.Replace(trackName, " ", "-", -1)
	artistName = strings.Replace(artistName, " ", "-", -1)
	var url = fmt.Sprintf(`http://api.musixmatch.com/ws/1.1/track.search?apikey=%s&page_size=1&q_track=%s&q_artist=%s`, API_KEY, trackName, artistName)
	res, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	response := tracks.Response{}
	json.Unmarshal(responseData, &response)
	return response.Message.Body.TrackList[0].Track
}
