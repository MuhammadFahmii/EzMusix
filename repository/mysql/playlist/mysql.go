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

func (repo *PlaylistRepo) Insert(playlistDomain playlist.Domain) (playlist.Domain, error) {
	rec := fromDomain(playlistDomain)
	if err := repo.DBConn.Create(&rec).Error; err != nil {
		return playlist.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *PlaylistRepo) Get(playlistDomain playlist.Domain) ([]playlist.Domain, error) {
	rec := []Playlist{}
	if err := repo.DBConn.Preload("Tracks").Find(&rec).Error; err != nil {
		return []playlist.Domain{}, err
	}
	var domainPlaylist []playlist.Domain
	for _, val := range rec {
		domainPlaylist = append(domainPlaylist, val.toDomain())
	}

	return domainPlaylist, nil
}

func (repo *PlaylistRepo) Delete(playlistDomain playlist.Domain) (playlist.Domain, error) {
	rec := fromDomain(playlistDomain)
	if err := repo.DBConn.Where("id = ?", playlistDomain.Id).Delete(&rec).Error; err != nil {
		return playlist.Domain{}, err
	}
	return rec.toDomain(), nil
}
