package tracks

import (
	"EzMusix/bussiness/tracks"

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

func (tracksRepo *TracksRepo) AddTrackPlaylist(detailPlaylist tracks.TrackPlaylist, track tracks.Domain) (tracks.Domain, error) {
	newTrack := FromDomain(track)
	if err := tracksRepo.DBConn.Model(&Playlist{Id: detailPlaylist.PlaylistId}).Association("Tracks").Append(&newTrack); err != nil {
		return tracks.Domain{}, err
	}
	return ToDomain(newTrack), nil
}

func (tracksRepo *TracksRepo) DeleteTrackPlaylist(playlistId, trackId int) (tracks.DeleteTrackPlaylist, error) {
	playlist := Playlist{}
	track := Track{}
	tracksRepo.DBConn.Select("name").Joins("LEFT JOIN detail_playlist ON detail_playlist.playlist_id = playlists.id").Find(&playlist)
	tracksRepo.DBConn.Select("track_name").Joins("LEFT JOIN detail_playlist ON detail_playlist.track_id = tracks.id").Find(&track)
	if err := tracksRepo.DBConn.Model(&Playlist{Id: playlistId}).Association("Tracks").Delete(&Track{Id: trackId}); err != nil {
		return tracks.DeleteTrackPlaylist{}, err
	}
	return tracks.DeleteTrackPlaylist{PlaylistName: playlist.Name, TrackName: track.TrackName}, nil
}
