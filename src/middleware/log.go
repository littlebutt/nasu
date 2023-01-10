package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/log"
	"time"
)

func LogRequired() gin.HandlerFunc {
	logger := log.Log.GetLogger()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUrl := c.Request.URL
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		params := ""
		if reqMethod == "POST" {
			params = c.Request.PostForm.Encode()
		}
		logger.Infof("[Nasu-handler]| %3d | %5v | %15s | %s | %s| %+v",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUrl,
			params,
		)
	}
}
