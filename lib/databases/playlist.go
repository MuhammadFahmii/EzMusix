package databases

import (
	"EzMusix/config"
	"EzMusix/lib/third_party"
	"EzMusix/models/playlists"
	"EzMusix/models/tracks"
)

func AddPlaylist(playlist *playlists.Playlist) (interface{}, error) {
	if err := config.DB.Debug().Create(&playlist).Error; err != nil {
		return nil, err
	}
	return &playlist, nil
}

func GetPlaylist(playlist *[]playlists.Playlist) (interface{}, error) {
	if err := config.DB.Preload("Tracks").Find(&playlist).Error; err != nil {
		return nil, err
	}
	return &playlist, nil
}

func DeletePlaylist(playlist *playlists.Playlist) (interface{}, error) {
	if err := config.DB.Where("id=?", playlist.Id).Delete(&playlist).Error; err != nil {
		return nil, err
	}
	return &playlist, nil
}

/*
DetailPlaylist Method
*/
func AddDetailPlaylist(detailPlaylist *playlists.DetailPlaylist) (interface{}, error) {
	newTrack := third_party.Search(detailPlaylist.TrackName, detailPlaylist.ArtistName)
	if err := config.DB.Model(&playlists.Playlist{Id: detailPlaylist.PlaylistId}).Association("Tracks").Append(&newTrack); err != nil {
		return nil, err
	}
	return &newTrack, nil
}

func DeleteDetailPlaylist(playlistId, trackId int) (interface{}, error) {
	arrResp := []int{playlistId, trackId}
	if err := config.DB.Model(&playlists.Playlist{Id: playlistId}).Association("Tracks").Delete(&tracks.Track{Id: trackId}); err != nil {
		return nil, err
	}
	return arrResp, nil
}
