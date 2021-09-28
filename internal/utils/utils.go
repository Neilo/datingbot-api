package utils

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/constant"
)

//GetReqIDFromContext return req-id from context
//if req id is empty return constant.EmptyReq
func GetReqIDFromContext(ctx context.Context) string {
	reqID, ok := ctx.Value(constant.ReqID).(string)
	if !ok {
		reqID = constant.EmptyReq
	}
	return reqID
}
