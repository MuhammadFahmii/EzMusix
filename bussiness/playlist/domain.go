package playlist

import "EzMusix/repository/mysql/tracks"

type Playlist struct {
	Id     int
	Name   string
	UserID int
	Tracks []tracks.Track `gorm:"many2many:detail_playlist"`
}
type Usecase interface {
	Get(playlist *Playlist) ([]*Playlist, error)
}

type Repository interface {
	Get(playlist *Playlist) ([]*Playlist, error)
}
