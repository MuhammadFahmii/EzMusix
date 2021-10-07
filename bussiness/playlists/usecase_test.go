package playlists_test

import (
	"EzMusix/bussiness/playlists"
	"EzMusix/bussiness/playlists/mocks"
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
	playlistsRepo.On("Insert", mock.Anything).Return(playlistDomain, nil)
	t.Run("Test Case 1 | Valid Data", func(t *testing.T) {
		_, err := playlistUseCase.Insert(playlistDomain)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}

func TestGet(t *testing.T) {
	testSetup()
	playlistsRepo.On("Get", mock.Anything).Return([]playlists.Domain{}, nil)
	t.Run("Test Case 1 | Valid Data", func(t *testing.T) {
		_, err := playlistUseCase.Get(playlistDomain)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}

func TestDelete(t *testing.T) {
	testSetup()
	playlistsRepo.On("Delete", mock.Anything).Return(playlistDomain, nil)
	t.Run("Test Case 2 | Data Found", func(t *testing.T) {
		_, err := playlistUseCase.Delete(playlistDomain)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}
