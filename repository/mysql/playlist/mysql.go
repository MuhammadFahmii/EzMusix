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

func (repo *PlaylistRepo) Get(playlistDomain *playlist.Domain) ([]playlist.Domain, error) {
	if err := repo.DBConn.Preload("Tracks").Find(&playlistDomain).Error; err != nil {
		return nil, err
	}
	pd := []playlist.Domain{}
	return pd, nil
}
