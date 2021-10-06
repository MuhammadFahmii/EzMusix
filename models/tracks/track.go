package tracks

type Track struct {
	Id         int    `gorm:"primarykey" json:"track_id" query:"track_id"`
	Name       string `json:"track_name" query:"track_name"`
	AlbumName  string `json:"album_name" query:"album_name"`
	ArtistName string `json:"artist_name" query:"artist_name"`
}
