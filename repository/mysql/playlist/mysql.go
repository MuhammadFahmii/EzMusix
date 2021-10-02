package playlist

import (
	"EzMusix/bussiness/playlist"

	"gorm.io/gorm"
)

type PlaylistRepo struct {
	DBConn *gorm.DB
}

func NewPlaylistRepo(db *gorm.DB) playlist.Repository {
	return &PlaylistRepo{
		DBConn: db,
	}
}

func (repo *PlaylistRepo) Get(playlistDomain *playlist.Playlist) ([]*playlist.Playlist, error) {
	if err := repo.DBConn.Preload("Tracks").Find(&playlistDomain).Error; err != nil {
		return []*playlist.Playlist{}, err
	}
	manyPlaylist := []*playlist.Playlist{}
	manyPlaylist = append(manyPlaylist, playlistDomain)
	return manyPlaylist, nil
}
