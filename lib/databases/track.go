package databases

import (
	"EzMusix/config"
	"EzMusix/models/tracks"
)

func GetTrack(track *[]tracks.Track) (interface{}, error) {
	if err := config.DB.Find(&track).Error; err != nil {
		return nil, err
	}
	return &track, nil
}

func DeleteTrack(track *tracks.Track) (interface{}, error) {
	if err := config.DB.Where("track_id = ?", track.Id).Delete(&track).Error; err != nil {
		return nil, err
	}
	return &track, nil
}
