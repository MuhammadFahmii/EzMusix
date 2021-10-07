package users

import (
	responseHandler "EzMusix/app/presenter"
	"EzMusix/app/presenter/users/request"
	"EzMusix/app/presenter/users/response"
	"EzMusix/bussiness/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	usersUC users.Usecase
}

func NewHandler(users users.Usecase) *Presenter {
	return &Presenter{
		usersUC: users,
	}
}

func (presenter *Presenter) Register(c echo.Context) error {
	reqUsers := request.Users{}
	c.Bind(&reqUsers)
	domain := reqUsers.ToDomain()
	res, err := presenter.usersUC.Register(domain)
	if err != nil {
		if err.Error() == "please fill all fields" || err.Error() == "duplicate entry" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusCreated, response.FromUsersRegister(res))
}

func (presenter *Presenter) Login(c echo.Context) error {
	reqUsers := request.Users{}
	c.Bind(&reqUsers)

	domain := reqUsers.ToDomain()
	res, err := presenter.usersUC.Login(domain)
	if err != nil {
		if err.Error() == "please fill all fields" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		if err.Error() == "record not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, response.FromUsersLogin(res))
}

func (presenter *Presenter) GetAllUsers(c echo.Context) error {
	usersDomain := users.Domain{}
	res, err := presenter.usersUC.GetAllUsers(usersDomain)
	if err != nil {
		if err.Error() == "not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	usersPlaylist := []response.UsersPlaylist{}
	for _, val := range res {
		usersPlaylist = append(usersPlaylist, response.ToUsersPlaylist(val))
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, usersPlaylist)
}
