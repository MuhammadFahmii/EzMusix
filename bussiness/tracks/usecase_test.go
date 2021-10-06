package tracks_test

import (
	"EzMusix/bussiness/tracks"
	"EzMusix/bussiness/tracks/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
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
	t.Run("Test Case 1 | Empty Data", func(t *testing.T) {
		tracksRepo.On("Get", mock.Anything, mock.Anything).Return(tracks.Domain{}, errors.New("Empty Data")).Once()
		_, err := tracksUsecase.Get(tracksDomain.TrackName, tracksDomain.ArtistName)
		assert.NotEqual(t, nil, err, "Data should not be empty")
	})
	t.Run("Test Case 2| Data Found", func(t *testing.T) {
		tracksRepo.On("Get", mock.Anything, mock.Anything).Return(tracks.Domain{}, nil).Once()
		tracksUsecase.Get(tracksDomain.TrackName, tracksDomain.ArtistName)
	})
}

func TestAddDetailPlaylist(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Internal Server Error", func(t *testing.T) {
		tracksRepo.On("AddDetailPlaylist", mock.Anything).Return(tracks.Domain{}, errors.New("Internal Server Error")).Once()
		_, err := tracksUsecase.AddDetailPlaylist(tracksDomain)
		assert.Equal(t, nil, err, "Add detail playlist not work")
	})
	t.Run("Test Case 2 | Success", func(t *testing.T) {
		tracksRepo.On("AddDetailPlaylist", mock.Anything).Return(tracks.Domain{}, nil).Once()
		tracksUsecase.AddDetailPlaylist(tracksDomain)
	})
}
func TestDeleteDetailPlaylist(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Internal Server Error", func(t *testing.T) {
		tracksRepo.On("DeleteDetailPlaylist", mock.Anything, mock.Anything).Return(tracks.DeleteTrackPlaylist{}, errors.New("Data Not Found")).Once()
		_, err := tracksUsecase.DeleteDetailPlaylist(tracksDomain.PlaylistId, 1)
		assert.Equal(t, nil, err, "Data not found")
	})
	t.Run("Test Case 2 | Success", func(t *testing.T) {
		tracksRepo.On("DeleteDetailPlaylist", mock.Anything, mock.Anything).Return(tracks.DeleteTrackPlaylist{}, nil).Once()
		tracksUsecase.DeleteDetailPlaylist(tracksDomain.PlaylistId, 1)
	})
}
