package playlists

import "EzMusix/models/tracks"

type Playlist struct {
	Id     int            `gorm:"primaryKey" json:"id"`
	Name   string         `json:"name" form:"name"`
	UserID int            `json:"user_id" form:"user_id"`
	Tracks []tracks.Track `gorm:"many2many:detail_playlist constraint:OnDelete:CASCADE"`
}
