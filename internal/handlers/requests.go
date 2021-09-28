package handlers

import (
	"github.com/brotherhood228/dating-bot-api/internal/errors"
	"github.com/brotherhood228/dating-bot-api/internal/model"
	"github.com/labstack/echo"
	"net/http"
)

//UserRequest
type UserRequest struct {
	User *model.User `json:"user,omitempty"`
	ID   *string     `json:"id,omitempty" query:"id"`
}

//ErrorResponse response bad answer from internal error
func ErrorResponse(c echo.Context, err error) error {
	answer, ok := err.(errors.Error)
	if !ok {
		return c.JSON(http.StatusBadRequest, errors.Error{MSG: err.Error()})
	}
	return c.JSON(answer.Status, answer)
}
