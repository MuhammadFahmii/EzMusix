package users_test

import (
	"EzMusix/app/middlewares"
	"EzMusix/bussiness/users"
	"EzMusix/bussiness/users/mocks"
	"fmt"

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
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(usersDomain, nil).Once()
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
	t.Run("Test Case 3 | Not Valid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(usersDomain, nil).Once()
		_, err := usersUsecase.Login(users.Domain{Username: "fahmi", Password: "123"})
		fmt.Println(err)
	})
}

func TestRegister(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(usersDomain)
	})
	t.Run("Test Case 2 | Not Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(users.Domain{Username: ""})
	})
	t.Run("Test Case 2 | Not Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(users.Domain{Username: "Fahmi", Password: ""})
	})
}
