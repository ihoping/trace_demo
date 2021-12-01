package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"trace_demo/pkg/common"
)

const clientIDHeader = "X-CLIENT-ID"
const clientIDCookieName = "client_id"

const traceIDHeader = "X-TRACE-ID"

func GetContext(ctx *gin.Context) context.Context {
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

	traceData := &common.TraceData{
		ClientID: clientID,
		TraceID:  traceID,
		Action:   action,
	}

	return common.GetContext(traceData)
}
