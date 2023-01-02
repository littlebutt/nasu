package middleware

import (
	"github.com/gin-gonic/gin"
	"nasu/src/misc"
	"time"
)

func LogRequired() gin.HandlerFunc {
	logger := misc.GetContextInstance().Logger
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUrl := c.Request.URL
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUrl,
		)
	}
}
