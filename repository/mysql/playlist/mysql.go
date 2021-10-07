package playlist

import (
	"EzMusix/bussiness/playlists"
	"errors"

	"gorm.io/gorm"
)

type PlaylistRepo struct {
	DBConn *gorm.DB
}

func NewPlaylistRepo(db *gorm.DB) playlists.Repository {
	return &PlaylistRepo{
		DBConn: db,
	}
}

func (repo *PlaylistRepo) Insert(playlistDomain playlists.Domain) (playlists.Domain, error) {
	rec := fromDomain(playlistDomain)
	if err := repo.DBConn.Create(&rec).Error; err != nil {
		return playlists.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *PlaylistRepo) Get(playlistDomain playlists.Domain) ([]playlists.Domain, error) {
	rec := []Playlist{}
	if err := repo.DBConn.Debug().Preload("Tracks").Find(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []playlists.Domain{}, errors.New("record not found")
		}
		return []playlists.Domain{}, err
	}
	var domainPlaylist []playlists.Domain
	for _, val := range rec {
		domainPlaylist = append(domainPlaylist, val.toDomain())
	}
	return domainPlaylist, nil
}

func (repo *PlaylistRepo) Delete(playlistDomain playlists.Domain) (playlists.Domain, error) {
	rec := fromDomain(playlistDomain)
	if err := repo.DBConn.Where("id=?", playlistDomain.Id).First(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return playlists.Domain{}, errors.New("record not found")
		}
	}
	repo.DBConn.Where("id = ?", playlistDomain.Id).Delete(&rec)
	return rec.toDomain(), nil
}
