package services

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/model"
	"github.com/brotherhood228/dating-bot-api/internal/utils"
	log "github.com/sirupsen/logrus"
)

//UserCreate service create user
func UserCreate(ctx context.Context, user *model.User, userRepo model.UserRepository) (string, error) {
	reqID := utils.GetReqIDFromContext(ctx)
	logger := log.WithFields(log.Fields{
		constant.ReqID: reqID,
	})
	logger.Debug("Enter in UserCreate service")

	return userRepo.Store(ctx, user)
}
