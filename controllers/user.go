package controllers

import (
	"EzMusix/lib/databases"
	"EzMusix/middlewares"
	"EzMusix/models/users"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUser(c echo.Context) error {
	user := users.User{}
	c.Bind(&user)
	if _, err := databases.AddUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":  "Success",
		"data": user,
	})
}
func LoginUser(c echo.Context) error {
	user := users.User{}
	c.Bind(&user)
	if _, err := databases.LoginUser(&user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, echo.Map{
				"msg": "Username / Password salah",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"msg": err.Error(),
		})
	}

	token, _ := middlewares.GenerateToken(user.Id)
	return c.JSON(http.StatusOK, echo.Map{
		"msg":   "Success",
		"data":  user,
		"token": token,
	})
}
func GetUser(c echo.Context) error {
	user := []users.User{}
	if _, err := databases.GetUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": user,
	})
}
