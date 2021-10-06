package tracks_test

import (
	"EzMusix/bussiness/tracks"
	"EzMusix/bussiness/tracks/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

var tracksRepo mocks.ThirdParty
var tracksDomain tracks.TrackPlaylist
var tracksUsecase tracks.Usecase

func testSetup() {
	tracksDomain = tracks.TrackPlaylist{
		PlaylistId: 1,
		ArtistName: "Pamungkas",
		TrackName:  "To The Bone",
	}
	tracksUsecase = tracks.NewTracksUsecase(&tracksRepo)
}

func TestGet(t *testing.T) {
	testSetup()
	t.Run("Test Case 2| Data Found", func(t *testing.T) {
		tracksRepo.On("Get", mock.Anything, mock.Anything).Return(tracks.Domain{}, nil).Once()
		tracksUsecase.Get(tracksDomain.TrackName, tracksDomain.ArtistName)
	})
}

func TestAddDetailPlaylist(t *testing.T) {
	testSetup()
	t.Run("Test Case 2 | Success", func(t *testing.T) {
		tracksRepo.On("AddDetailPlaylist", mock.Anything).Return(tracks.Domain{}, nil).Once()
		tracksUsecase.AddDetailPlaylist(tracksDomain)
	})
}
func TestDeleteDetailPlaylist(t *testing.T) {
	testSetup()
	t.Run("Test Case 2 | Success", func(t *testing.T) {
		tracksRepo.On("DeleteDetailPlaylist", mock.Anything, mock.Anything).Return(tracks.DeleteTrackPlaylist{}, nil).Once()
		tracksUsecase.DeleteDetailPlaylist(tracksDomain.PlaylistId, 1)
	})
}
