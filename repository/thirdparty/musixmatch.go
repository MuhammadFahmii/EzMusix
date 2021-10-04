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

func (tracksRepo *TracksRepo) Get(trackName, artistName string) (tracks.Domain, error) {
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
	return response.toDomain(), nil
}

func (tracksRepo *TracksRepo) AddDetailPlaylist(detailPlaylist tracks.TrackPlaylist) (tracks.Domain, error) {
	newTrack, err := tracksRepo.Get(detailPlaylist.TrackName, detailPlaylist.ArtistName)
	if err != nil {
		return tracks.Domain{}, err
	}
	if err := tracksRepo.DBConn.Model(&Playlist{Id: detailPlaylist.PlaylistId}).Association("Tracks").Append(&newTrack); err != nil {
		return tracks.Domain{}, err
	}
	return newTrack, nil
}

func (tracksRepo *TracksRepo) DeleteDetailPlaylist(playlistId, trackId int) (tracks.DeleteTrackPlaylist, error) {
	playlist := Playlist{}
	track := Track{}
	tracksRepo.DBConn.Debug().Select("name").Joins("LEFT JOIN detail_playlist ON detail_playlist.playlist_id = playlists.id").Find(&playlist)
	tracksRepo.DBConn.Debug().Select("name").Joins("LEFT JOIN detail_playlist ON detail_playlist.track_id = tracks.id").Find(&track)
	if err := tracksRepo.DBConn.Model(&Playlist{Id: playlistId}).Association("Tracks").Delete(&tracks.Domain{Id: trackId}); err != nil {
		return tracks.DeleteTrackPlaylist{}, err
	}
	return tracks.DeleteTrackPlaylist{PlaylistName: playlist.Name, TrackName: track.Name}, nil
}
