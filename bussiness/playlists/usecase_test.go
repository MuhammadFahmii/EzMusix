package playlists_test

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/playlists/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var playlistsRepo mocks.Repository
var playlistDomain playlists.Domain
var playlistUseCase playlists.Usecase

func testSetup() {
	playlistDomain = playlists.Domain{
		Id:     1,
		Name:   "Happy",
		UserID: 1,
	}
	playlistUseCase = playlists.NewPlaylistUsecase(&playlistsRepo)
}

func TestInsert(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Data not valid", func(t *testing.T) {
		playlistsRepo.On("Insert", mock.Anything).Return(playlists.Domain{}, errors.New("Data not valid")).Once()
		_, err := playlistUseCase.Insert(playlistDomain)
		assert.NotEqual(t, nil, err, "Please input valid data")
	})
	t.Run("Test Case 2 | Valid Data", func(t *testing.T) {
		playlistsRepo.On("Insert", mock.Anything).Return(playlistDomain, nil).Once()
		playlistUseCase.Insert(playlistDomain)
	})
}

func TestGet(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Data Empty", func(t *testing.T) {
		playlistsRepo.On("Get", mock.Anything).Return([]playlists.Domain{}, errors.New("Data empty")).Once()
		_, err := playlistUseCase.Get(playlistDomain)
		assert.NotEqual(t, nil, err, "Data should not be empty")
	})
	t.Run("Test Case 2 | Data Found", func(t *testing.T) {
		playlistsRepo.On("Get", mock.Anything).Return([]playlists.Domain{}, nil).Once()
		playlistUseCase.Get(playlistDomain)
	})
}

func TestDelete(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Data not found", func(t *testing.T) {
		playlistsRepo.On("Delete", mock.Anything).Return(playlists.Domain{}, errors.New("Data not found")).Once()
		_, err := playlistUseCase.Delete(playlistDomain)
		assert.NotEqual(t, nil, err, "Please input valid data")
	})
	t.Run("Test Case 2 | Data Found", func(t *testing.T) {
		playlistsRepo.On("Delete", mock.Anything).Return(playlistDomain, nil).Once()
		playlistUseCase.Delete(playlistDomain)
	})
}
