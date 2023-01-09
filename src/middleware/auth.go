package middleware

import (
	"github.com/gin-gonic/gin"
	"nasu/src/context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		res := strings.Split(authorization, "+")
		if len(res) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		targetPassword := res[0]
		targetTimestamp, err := strconv.Atoi(res[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		password := context.NasuContext.Password
		// redirect if timeout
		if targetPassword != password || (time.Now().Unix()-int64(targetTimestamp) > context.NasuContext.TokenTTL*60*60) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
