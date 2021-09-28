package services

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/model"
	"github.com/brotherhood228/dating-bot-api/internal/utils"
	log "github.com/sirupsen/logrus"
)

//UserUpdate service create user
func UserUpdate(ctx context.Context, user *model.User, id *string, userRepo model.UserRepository) error {
	reqID := utils.GetReqIDFromContext(ctx)
	logger := log.WithFields(log.Fields{
		constant.ReqID: reqID,
	})
	logger.Debug("Enter in UserUpdate service")

	return userRepo.Update(ctx, id, user)
}
