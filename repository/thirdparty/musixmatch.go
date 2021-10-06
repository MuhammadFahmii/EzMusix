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

func NewTracksRepo(conn *gorm.DB) tracks.ThirdParty {
	return &TracksRepo{
		DBConn: conn,
	}
}

func (tracksRepo *TracksRepo) Get(trackName, artistName string) (tracks.Track, error) {
	const API_KEY = "4b69788a296733b380069f770a174a89"
	trackName = strings.Replace(trackName, " ", "-", -1)
	artistName = strings.Replace(artistName, " ", "-", -1)
	var url = fmt.Sprintf(`http://api.musixmatch.com/ws/1.1/track.search?apikey=%s&page_size=1&q_track=%s&q_artist=%s`, API_KEY, trackName, artistName)
	res, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	response := Response{}
	if err := json.Unmarshal(responseData, &response); err != nil {
		return tracks.Track{}, err
	}
	return response.toDomain(), nil
}

func (tracksRepo *TracksRepo) AddDetailPlaylist(detailPlaylist tracks.DetailPlaylist) (tracks.Track, error) {
	newTrack, err := tracksRepo.Get(detailPlaylist.TrackName, detailPlaylist.ArtistName)
	if err != nil {
		return tracks.Track{}, err
	}
	if err := tracksRepo.DBConn.Model(&Playlist{Id: detailPlaylist.PlaylistId}).Association("Tracks").Append(&newTrack); err != nil {
		return tracks.Track{}, err
	}
	return newTrack, nil
}

func (tracksRepo *TracksRepo) DeleteDetailPlaylist(playlistId, trackId int) (tracks.DetailPlaylist, error) {
	detailPlaylist := Playlist{}
	tracksRepo.DBConn.Debug().Select("name").Joins("left join detail_playlist on detail_playlist.playlist_id = playlists.id").Find(&detailPlaylist)
	fmt.Println(detailPlaylist)
	if err := tracksRepo.DBConn.Model(&Playlist{Id: playlistId}).Association("Tracks").Delete(&tracks.Track{Id: trackId}); err != nil {
		return tracks.DetailPlaylist{}, err
	}
	return tracks.DetailPlaylist{}, nil
}
