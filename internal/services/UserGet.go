package services

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/model"
	"github.com/brotherhood228/dating-bot-api/internal/utils"
	log "github.com/sirupsen/logrus"
)

//UserGet service return user by id
func UserGet(ctx context.Context, id *string, userRepo model.UserRepository) (*model.User, error) {
	reqID := utils.GetReqIDFromContext(ctx)
	logger := log.WithFields(log.Fields{
		constant.ReqID: reqID,
		"user_id":      id,
	})
	logger.Debug("Enter in UserGet service")
	return userRepo.Find(ctx, id)
}
