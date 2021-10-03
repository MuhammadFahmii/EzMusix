package users_test

import (
	"EzMusix/bussiness/users"
	"EzMusix/bussiness/users/mocks"

	"testing"

	"github.com/stretchr/testify/mock"
)

var usersRepository mocks.Repository
var usersUsecase users.Usecase
var usersDomain users.Users

func testSetup() {
	usersUsecase = users.NewUserUsecase(&usersRepository)
	usersDomain = users.Users{
		Username: "fahmi",
		Password: "123",
	}
}

func TestLogin(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		usersRepository.On("Login", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Login(usersDomain)
	})
}

func TestRegister(t *testing.T) {
	testSetup()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()
		usersUsecase.Register(usersDomain)
	})
}
