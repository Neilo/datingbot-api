package handlers

import (
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/errors"
	"github.com/brotherhood228/dating-bot-api/internal/services"
	"github.com/brotherhood228/dating-bot-api/internal/stores"
	"github.com/brotherhood228/dating-bot-api/pkg/mongo"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const updatePhoto = "update_photo"

const getPhoto = "get_photo"

//UpdatePhoto controller for update Photo
func UpdatePhoto(c echo.Context) error {
	ctx, reqID := prepareReqID(c)

	fileForm, err := c.FormFile("file")
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v when read from part form", updatePhoto)
		return ErrorResponse(c, err)
	}

	id := c.FormValue("id")

	file, err := fileForm.Open()
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v when open file", updatePhoto)
		return ErrorResponse(c, err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v ", updatePhoto)
		return ErrorResponse(c, err)
	}

	//prepare stores
	mongoSession, err := mongo.GetDB()

	photoStore := stores.InitPhoto(mongoSession.Session.Copy())

	err = services.UpdatePhoto(ctx, id, data, photoStore)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v ", updatePhoto)
		return ErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, photoResponse{Write: true})
}

//GetPhoto controller for get Photo
func GetPhoto(c echo.Context) error {
	ctx, reqID := prepareReqID(c)

	id := c.QueryParam("id")

	if len(id) == 0 {
		log.WithFields(log.Fields{
			"Error":        "Empty id",
			constant.ReqID: reqID,
		}).Debugf("Error in %v", getPhoto)
		return ErrorResponse(c, errors.BadRequest)
	}

	//prepare stores
	mongoSession, err := mongo.GetDB()

	photoStore := stores.InitPhoto(mongoSession.Session.Copy())

	photo, err := services.GetPhoto(ctx, id, photoStore)
	if err != nil {
		log.WithFields(log.Fields{
			"Error":        err,
			constant.ReqID: reqID,
		}).Debugf("Error in %v", getPhoto)
		return ErrorResponse(c, err)
	}

	return c.Blob(http.StatusOK, "image/jpeg", photo)
}

type photoResponse struct {
	Write bool `json:"status"`
}
