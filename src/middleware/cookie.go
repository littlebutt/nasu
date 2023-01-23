package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CookieRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		targetPassword := strings.Split(token, " ")[0]
		if len(strings.Split(token, " ")) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		targetTimestamp, err := strconv.Atoi(strings.Split(token, " ")[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		password := context.NasuContext.Password
		if targetPassword != password || (time.Now().Unix()-int64(targetTimestamp) > context.NasuContext.TokenTTL*60*60) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
