package thirdparty

import (
	"EzMusix/bussiness/tracks"
)

type Response struct {
	Message Message `json:"message"`
}
type Message struct {
	Body Body `json:"body"`
}
type Body struct {
	TrackList []TrackList `json:"track_list"`
}
type TrackList struct {
	Track Track `json:"track"`
}
type Track struct {
	Id         int    `json:"track_id"`
	Name       string `json:"track_name"`
	ArtistName string `json:"artist_name"`
	AlbumName  string `json:"album_name"`
}

func (pl *Response) toDomain() tracks.Domain {
	return tracks.Domain{
		Id:         pl.Message.Body.TrackList[0].Track.Id,
		Name:       pl.Message.Body.TrackList[0].Track.Name,
		ArtistName: pl.Message.Body.TrackList[0].Track.ArtistName,
		AlbumName:  pl.Message.Body.TrackList[0].Track.AlbumName,
	}
}
