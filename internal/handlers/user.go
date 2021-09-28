package handlers

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/errors"
	"github.com/brotherhood228/dating-bot-api/internal/services"
	"github.com/brotherhood228/dating-bot-api/internal/stores"
	"github.com/brotherhood228/dating-bot-api/pkg/metric"
	"github.com/brotherhood228/dating-bot-api/pkg/postgree"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"net/http"
)

const (
	userHandler     = "user_handler"
	userPickHandler = "user_pick_handler"
)

var newUsersMetric = metric.MustRegisterCounter("new_users", "Counter for new user")

var newUsersErrorMetric = metric.MustRegisterCounter("new_users_error", "Errors when create new user")

var SuccessResponseMetric = metric.MustRegisterCounterVec("success_responses_api", "Count for success responses", []string{"operations"})

//UserHandler is handler for user crud
func UserHandler(c echo.Context) error {
	ctx, reqID := prepareReqID(c)

	var req UserRequest
	err := c.Bind(&req)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v when bind request", userHandler)
		err := errors.BadRequest.New(err)
		return ErrorResponse(c, err)
	}

	db, err := postgree.GetDB()
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v when get pg instance", userHandler)
		err := errors.DataBaseError.New(err)
		return ErrorResponse(c, err)
	}
	repo := stores.InitUserStore(db)

	switch c.Request().Method {
	case http.MethodGet:
		user, err := services.UserGet(ctx, req.ID, repo)
		if err != nil {
			return ErrorResponse(c, err)
		}
		return SuccessResponse(c, user, req.ID, http.StatusOK)
	case http.MethodPost:
		id, err := services.UserCreate(ctx, req.User, repo)
		if err != nil {
			newUsersErrorMetric.Inc()
			return ErrorResponse(c, err)
		}
		newUsersMetric.Inc()
		return SuccessResponse(c, nil, &id, http.StatusCreated)
	case http.MethodPut:
		err := services.UserUpdate(ctx, req.User, req.ID, repo)
		if err != nil {
			return ErrorResponse(c, err)
		}
		return SuccessResponse(c, req.User, req.ID, http.StatusOK)
	}
	return ErrorResponse(c, errors.BadMethod)
}

func prepareReqID(c echo.Context) (context.Context, string) {
	reqID, ok := c.Get(constant.ReqID).(string)
	if len(reqID) == 0 || !ok {
		reqIDNew, err := shortid.Generate()
		if err != nil {
			log.Debug("Bad generate req-id in UserHandler")
			reqIDNew = constant.EmptyReq
		}
		reqID = reqIDNew
	}
	ctx := context.WithValue(c.Request().Context(), constant.ReqID, reqID)
	return ctx, reqID
}

//UserPick ...
func UserPick(c echo.Context) error {
	ctx, reqID := prepareReqID(c)

	var req UserRequest

	err := c.Bind(&req)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v when bind request", userPickHandler)
		err := errors.BadRequest.New(err)
		return ErrorResponse(c, err)
	}

	logger := log.WithFields(log.Fields{
		constant.ReqID: reqID,
		"payload":      req,
	})

	logger.Debug(userPickHandler)

	db, err := postgree.GetDB()
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v when get pg instance", userPickHandler)
		err := errors.DataBaseError.New(err)
		return ErrorResponse(c, err)
	}
	repo := stores.InitUserStore(db)

	user, err := services.UserPick(ctx, repo, req.ID)
	if err != nil {
		return ErrorResponse(c, err)
	}
	return SuccessResponse(c, user, nil, http.StatusOK)
}
