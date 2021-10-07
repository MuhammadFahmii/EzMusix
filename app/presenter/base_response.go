package presenter

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewSuccessResponse(c echo.Context, status int, param interface{}) error {
	response := SuccessResponse{}
	response.Status = status
	response.Message = "successful"
	response.Data = param

	return c.JSON(status, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := ErrorResponse{}
	response.Status = status
	response.Message = err.Error()
	return c.JSON(status, response)
}
