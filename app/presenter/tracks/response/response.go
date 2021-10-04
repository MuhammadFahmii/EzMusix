package response

import "EzMusix/bussiness/tracks"

type DeleteTrackPlaylist struct {
	PlaylistName string
	TrackName    string
}

func FromDomain(dtp tracks.DeleteTrackPlaylist) DeleteTrackPlaylist {
	return DeleteTrackPlaylist{
		PlaylistName: dtp.PlaylistName,
		TrackName:    dtp.TrackName,
	}
}
