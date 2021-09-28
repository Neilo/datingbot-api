package handlers

import (
	"github.com/brotherhood228/dating-bot-api/internal/model"
	"github.com/labstack/echo"
	"net/http"
)

//UserResponse answer for a user Handler
type UserResponse struct {
	User *model.User `json:"user,omitempty"`
	ID   *string     `json:"id,omitempty"`
}

func SuccessResponse(c echo.Context, user *model.User, id *string, status int) error {
	SuccessResponseMetric.WithLabelValues(c.Request().Method).Inc()
	if status == 0 {
		status = http.StatusOK
	}

	r := UserResponse{
		User: user,
		ID:   id,
	}

	return c.JSON(status, r)
}
