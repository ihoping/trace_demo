package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
	"trace_demo/pkg/common"
	"trace_demo/pkg/server"
)

const defaultMemory = 32 << 20

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := server.GetContext(c)

		logger, err := common.GetLogger(ctx)
		if err != nil {
			return
		}

		start := time.Now()

		method := c.Request.Method

		var logReqBody interface{}
		reqContentType := c.ContentType()
		if strings.HasPrefix(reqContentType, "x-www-form-urlencoded") {
			err := c.Request.ParseForm()
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			result := map[string]string{}
			for key, vals := range c.Request.PostForm {
				result[key] = vals[0]
			}
			logReqBody, err = json.Marshal(result)
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			logReqBody = fmt.Sprintf("Form(%s)", logReqBody)
		} else if strings.HasPrefix(reqContentType, "multipart/form-data") {
			err := c.Request.ParseMultipartForm(defaultMemory)
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			result := map[string]string{}
			for key, vals := range c.Request.MultipartForm.Value {
				result[key] = vals[0]
			}
			for key, files := range c.Request.MultipartForm.File {
				file := files[0]
				contentType := file.Header.Get("Content-Type")
				result[key] = fmt.Sprintf("%s(%s)[%d]", file.Filename, contentType, file.Size)
			}
			logReqBody, err = json.Marshal(result)
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			logReqBody = fmt.Sprintf("MultipartForm(%s)", logReqBody)
		} else if strings.HasPrefix(reqContentType, "application/json") {
			body, _ := io.ReadAll(c.Request.Body)
			err := c.Request.Body.Close()
			if err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			logReqBody = fmt.Sprintf("JSON(%s)", body)
		} else {
			if c.Request.ContentLength == 0 {
				logReqBody = "[EmptyData]"
			} else {
				logReqBody = "[BinaryData]"
			}
		}

		// Process request
		c.Next()

		end := time.Now()
		latency := end.Sub(start).Seconds()

		logger.Info(fmt.Sprintf("[gin] %s %fs %s", method, latency, logReqBody))
	}
}
