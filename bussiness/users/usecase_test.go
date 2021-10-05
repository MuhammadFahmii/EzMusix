package users_test

import (
	"EzMusix/app/middlewares"
	"EzMusix/bussiness/users"
	"EzMusix/bussiness/users/mocks"
	"errors"

	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var usersRepository mocks.Repository
var usersUsecase users.Usecase
var usersDomain users.Domain
var configJWT middlewares.ConfigJWT

func testSetup() {
	configJWT = middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("secret"),
		ExpiresDuration: viper.GetInt("expired"),
	}
	usersUsecase = users.NewUserUsecase(&usersRepository, &configJWT)
	usersDomain = users.Domain{
		Username: "fahmi",
		Password: "1234",
	}

}

func TestLogin(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Data Empty", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(users.Domain{}, errors.New("Data Empty")).Once()
		usersUsecase.Login(usersDomain)
	})
	t.Run("Test Case 2 | Not Valid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Login(users.Domain{Username: ""})
	})
	t.Run("Test Case 3 | Not Valid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Login(users.Domain{Username: "Fahmi", Password: ""})
	})
	t.Run("Test Case 4 | Valid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Login(usersDomain)
	})
}

func TestRegister(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Internal Server Error", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, errors.New("internal server error")).Once()
		_, err := usersUsecase.Register(usersDomain)
		assert.Equal(t, err, nil, "Register should work")
	})
	t.Run("Test Case 2 | Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(usersDomain)
	})
	t.Run("Test Case 3 | Not Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(users.Domain{Username: ""})
	})
	t.Run("Test Case 4 | Not Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(users.Domain{Username: "Fahmi", Password: ""})
	})
}

func TestGetAllUsers(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Data Empty", func(t *testing.T) {
		usersRepository.On("GetAllUsers", mock.Anything).Return(nil, errors.New("sss")).Once()
		usersUsecase.GetAllUsers(usersDomain)
	})
	t.Run("Test Case 2| Get All Data", func(t *testing.T) {
		usersRepository.On("GetAllUsers", mock.Anything).Return([]users.Domain{}, nil).Once()
		usersUsecase.GetAllUsers(usersDomain)
	})
}
