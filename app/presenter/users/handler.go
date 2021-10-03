package users

import (
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
	if err := c.Bind(&reqUsers); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	domain := reqUsers.ToDomain()
	res, err := presenter.usersUC.Register(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response.FromUsersRegister(res))
}

func (presenter *Presenter) Login(c echo.Context) error {
	reqUsers := request.Users{}
	if err := c.Bind(&reqUsers); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	domain := reqUsers.ToDomain()
	res, err := presenter.usersUC.Login(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, response.FromUsersLogin(res))
}
