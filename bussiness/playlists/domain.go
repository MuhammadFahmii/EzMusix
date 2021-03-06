package playlists

import (
	"EzMusix/bussiness/tracks"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Tracks    []tracks.Domain
}

type Usecase interface {
	Insert(playlist Domain) (Domain, error)
	Get(playlist Domain) ([]Domain, error)
	Delete(playlist Domain) (Domain, error)
}

type Repository interface {
	Insert(playlist Domain) (Domain, error)
	Get(playlist Domain) ([]Domain, error)
	Delete(playlist Domain) (Domain, error)
}
