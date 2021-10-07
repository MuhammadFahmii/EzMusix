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
	Id            int    `json:"track_id"`
	TrackName     string `json:"track_name"`
	ArtistName    string `json:"artist_name"`
	AlbumName     string `json:"album_name"`
	TrackRating   int    `json:"track_rating"`
	TrackShareUrl string `json:"track_share_url"`
}

func (pl *Response) toDomain() tracks.Domain {
	return tracks.Domain{
		Id:            pl.Message.Body.TrackList[0].Track.Id,
		TrackName:     pl.Message.Body.TrackList[0].Track.TrackName,
		ArtistName:    pl.Message.Body.TrackList[0].Track.ArtistName,
		AlbumName:     pl.Message.Body.TrackList[0].Track.AlbumName,
		TrackRating:   pl.Message.Body.TrackList[0].Track.TrackRating,
		TrackShareUrl: pl.Message.Body.TrackList[0].Track.TrackShareUrl,
	}
}
