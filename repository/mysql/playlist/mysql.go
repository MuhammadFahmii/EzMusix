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

func (repo *PlaylistRepo) Insert(playlistDomain playlist.Playlist) (playlist.Playlist, error) {
	rec := fromDomain(playlistDomain)
	if err := repo.DBConn.Create(&rec).Error; err != nil {
		return playlist.Playlist{}, err
	}
	return rec.toDomain(), nil
}

func (repo *PlaylistRepo) Get(playlistDomain playlist.Playlist) ([]playlist.Playlist, error) {
	rec := []Playlist{}
	if err := repo.DBConn.Preload("Tracks").Find(&rec).Error; err != nil {
		return []playlist.Playlist{}, err
	}
	var domainPlaylist []playlist.Playlist
	for _, val := range rec {
		domainPlaylist = append(domainPlaylist, val.toDomain())
	}
	return domainPlaylist, nil
}

func (repo *PlaylistRepo) Delete(playlistDomain playlist.Playlist) (playlist.Playlist, error) {
	rec := fromDomain(playlistDomain)
	if err := repo.DBConn.Where("id = ?", playlistDomain.Id).Delete(&rec).Error; err != nil {
		return playlist.Playlist{}, err
	}
	return rec.toDomain(), nil
}
