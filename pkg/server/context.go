package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"trace_demo/pkg/common"
)

const clientIDHeader = "X-CLIENT-ID"
const clientIDCookieName = "client_id"

const traceIDHeader = "X-TRACE-ID"

const contextKey = "context"

func GetContext(ctx *gin.Context) context.Context {
	if exist, ok := ctx.Value(contextKey).(context.Context); ok {
		return exist
	}

	clientID := ctx.Request.Header.Get(clientIDHeader)
	if clientID == "" {
		clientID, _ = ctx.Cookie(clientIDCookieName)
		if clientID == "" {
			clientID = common.RandStringRunes(16)
			ctx.SetCookie(clientIDCookieName, clientID, 0, "/", "", true, true)
		}
	}

	traceID := ctx.Request.Header.Get(traceIDHeader)
	if traceID == "" {
		traceID = common.RandStringRunes(16)
	}

	action := ctx.Request.RequestURI

	logData := common.LogData{
		ClientID: clientID,
		TraceID:  traceID,
		Action:   action,
	}

	ctx2 := common.GetContext(logData)
	ctx.Set(contextKey, ctx2)
	return ctx2
}
