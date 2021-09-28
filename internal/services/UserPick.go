package services

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
	"github.com/brotherhood228/dating-bot-api/internal/model"
	log "github.com/sirupsen/logrus"
)

//UserPick pick user by id
//pick random user if subject doesn't like him
func UserPick(ctx context.Context, userStore model.UserRepository, id *string) (*model.User, error) {
	reqId, ok := ctx.Value(constant.ReqID).(string)
	if !ok {
		reqId = constant.EmptyReq
	}

	logger := log.WithFields(log.Fields{
		constant.ReqID: reqId,
		"id":           id,
		"service":      "user_pick",
	})

	logger.Debug("Enter in service")

	select {
	case <-ctx.Done():
		logger.Debug("Cancel service")
		return nil, nil
	default:
		return userStore.Pick(ctx, id)
	}
}
