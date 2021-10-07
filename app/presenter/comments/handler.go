package comments

import (
	responseHandler "EzMusix/app/presenter"
	"EzMusix/app/presenter/comments/request"
	"EzMusix/app/presenter/comments/response"
	"EzMusix/bussiness/comments"
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	commentsUC comments.Usecase
}

func NewHandler(pl comments.Usecase) *Presenter {
	return &Presenter{
		commentsUC: pl,
	}
}

func (presenter *Presenter) Insert(c echo.Context) error {
	reqcomments := request.Comments{}
	c.Bind(&reqcomments)
	domain := request.ToDomain(reqcomments)
	res, err := presenter.commentsUC.Insert(domain)
	resFromDomain := response.ToAddComments(res)
	if err != nil {
		if err.Error() == "please fill all fields" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusCreated, resFromDomain)
}

func (presenter *Presenter) Get(c echo.Context) error {
	reqcomments := request.Comments{}
	domain := request.ToDomain(reqcomments)
	res, err := presenter.commentsUC.Get(domain)
	if err != nil {
		if err.Error() == "not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	resFromDomain := []response.Comments{}
	for _, val := range res {
		resFromDomain = append(resFromDomain, response.FromDomain(val))
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, resFromDomain)
}

func (presenter *Presenter) Delete(c echo.Context) error {
	reqcomments := request.DeleteComments{}
	c.Bind(&reqcomments)
	domain := request.DeleteToDomain(reqcomments)
	res, err := presenter.commentsUC.Delete(domain)
	fmt.Println(res)
	if err != nil {
		if err.Error() == "record not found" {
			return responseHandler.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return responseHandler.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return responseHandler.NewSuccessResponse(c, http.StatusOK, response.FromDomainDelete(res))
}
