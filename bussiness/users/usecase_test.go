package users_test

import (
	"EzMusix/app/middlewares"
	"EzMusix/bussiness/users"
	"EzMusix/bussiness/users/mocks"

	"testing"

	"github.com/spf13/viper"
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
	usersRepository.On("Login", mock.Anything).Return(usersDomain, nil)
	t.Run("Test Case 2 | Not Valid Login", func(t *testing.T) {
		usersUsecase.Login(users.Domain{Username: ""})
	})
	t.Run("Test Case 3 | Not Valid Login", func(t *testing.T) {
		usersUsecase.Login(users.Domain{Username: "Fahmi", Password: ""})
	})
	t.Run("Test Case 4 | Valid Login", func(t *testing.T) {
		usersUsecase.Login(usersDomain)
	})
}

func TestRegister(t *testing.T) {
	testSetup()
	usersRepository.On("Register", mock.Anything).Return(usersDomain, nil)
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		usersUsecase.Register(usersDomain)
	})
	t.Run("Test Case 2| Not Valid Register", func(t *testing.T) {
		usersUsecase.Register(users.Domain{Username: ""})
	})
	t.Run("Test Case 3 | Not Valid Register", func(t *testing.T) {
		usersUsecase.Register(users.Domain{Username: "Fahmi", Password: ""})
	})
}

func TestGetAllUsers(t *testing.T) {
	testSetup()
	usersRepository.On("GetAllUsers", mock.Anything).Return([]users.Domain{}, nil)
	t.Run("Test Case 1 | Data Empty", func(t *testing.T) {
		usersUsecase.GetAllUsers(usersDomain)
	})
	t.Run("Test Case 2| Get All Data", func(t *testing.T) {
		usersUsecase.GetAllUsers(usersDomain)
	})
}
