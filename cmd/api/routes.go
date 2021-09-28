package main

import (
	"github.com/brotherhood228/dating-bot-api/internal/handlers"
	"github.com/brotherhood228/dating-bot-api/pkg/router"
	"net/http"
)

var routes = []router.Router{
	{
		Path:    "/user/photo",
		Method:  http.MethodPost,
		Handler: handlers.UpdatePhoto,
	},
	{
		Path:    "/user/photo",
		Method:  http.MethodGet,
		Handler: handlers.GetPhoto,
	},
	{
		Path:    "/user",
		Method:  router.AnyMethods,
		Handler: handlers.UserHandler,
	},
	{
		Path:    "/user/pick",
		Method:  http.MethodPost,
		Handler: handlers.UserPick,
	},
}
