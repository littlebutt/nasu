package middleware

import (
	"github.com/gin-gonic/gin"
	"nasu/src/misc"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AuthRequired() gin.HandlerFunc {
	password := misc.GetContextInstance().Password
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
		// redirect if timeout (3 * 60 * 60 s)
		// TODO: customize this value
		if targetPassword != password || (time.Now().Unix()-int64(targetTimestamp) > 3*60*60) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
