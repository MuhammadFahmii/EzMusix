package presenter

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, status int, param interface{}) error {
	response := BaseResponse{}
	response.Status = status
	response.Message = "successful"
	response.Data = param

	return c.JSON(status, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Status = status
	fmt.Println(err.Error())
	response.Message = err.Error()

	return c.JSON(status, response)
}
