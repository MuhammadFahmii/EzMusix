package playlist

import (
	"EzMusix/bussiness/tracks"
)

type Playlist struct {
	Id     int
	Name   string
	UserID int
	Tracks []tracks.Track
}

type Usecase interface {
	Insert(playlist Playlist) (Playlist, error)
	Get(playlist Playlist) ([]Playlist, error)
	Delete(playlist Playlist) (Playlist, error)
}

type Repository interface {
	Insert(playlist Playlist) (Playlist, error)
	Get(playlist Playlist) ([]Playlist, error)
	Delete(playlist Playlist) (Playlist, error)
}
