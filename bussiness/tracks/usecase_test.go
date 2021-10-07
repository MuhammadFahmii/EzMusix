package tracks_test

import (
	"EzMusix/bussiness/tracks"
	"EzMusix/bussiness/tracks/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var thirdParty mocks.ThirdParty
var tracksRepo mocks.Repository
var tracksDomain tracks.Domain
var trackPlaylist tracks.TrackPlaylist
var DeleteTrackPlaylist tracks.DeleteTrackPlaylist
var tracksUsecase tracks.Usecase

func testSetup() {
	trackPlaylist = tracks.TrackPlaylist{
		PlaylistId: 1,
		ArtistName: "Pamungkas",
		TrackName:  "To The Bone",
	}
	tracksUsecase = tracks.NewTracksUsecase(&tracksRepo, &thirdParty)
}

func TestGet(t *testing.T) {
	testSetup()
	thirdParty.On("Get", mock.Anything, mock.Anything).Return(tracksDomain, nil)
	t.Run("Test Case 1| Data Found", func(t *testing.T) {
		_, err := tracksUsecase.Get(trackPlaylist.TrackName, trackPlaylist.ArtistName)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}

func TestAddTrackPlaylist(t *testing.T) {
	thirdParty.On("Get", mock.Anything, mock.Anything).Return(tracksDomain, nil)
	tracksRepo.On("AddTrackPlaylist", mock.Anything, mock.Anything).Return(tracksDomain, nil)
	t.Run("Test Case 1| Data Found", func(t *testing.T) {
		tracksUsecase.Get(trackPlaylist.TrackName, trackPlaylist.ArtistName)
		_, err := tracksUsecase.AddTrackPlaylist(trackPlaylist)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}

func TestDeleteTrackPlaylist(t *testing.T) {
	testSetup()
	tracksRepo.On("DeleteTrackPlaylist", mock.Anything, mock.Anything).Return(DeleteTrackPlaylist, nil)
	t.Run("Test Case 1| Success", func(t *testing.T) {
		_, err := tracksUsecase.DeleteTrackPlaylist(trackPlaylist.PlaylistId, 1)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})

}
