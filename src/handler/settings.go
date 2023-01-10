package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/service"
	"net/http"
	"strconv"
)

func HandleChangeHashPrefix(c *gin.Context) {
	hashPrefix, err := strconv.Atoi(c.PostForm("hashPrefix"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": err.Error(),
		})
		return
	}
	res := service.ChangeHashPrefix(hashPrefix)
	c.JSON(http.StatusOK, gin.H{
		"success": res,
	})
}

func HandleChangeMaxFileSize(c *gin.Context) {
	size, err := strconv.Atoi(c.PostForm("size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": err.Error(),
		})
		return
	}
	res := service.ChangeMaxFileSize(size)
	c.JSON(http.StatusOK, gin.H{
		"success": res,
	})
}

func HandleChangeTokenTtl(c *gin.Context) {
	tokenTtlStr := c.PostForm("tokenTtl")
	tokenTtl, err := strconv.Atoi(tokenTtlStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": err.Error(),
		})
	}
	res := service.ChangeTokenTtl(int64(tokenTtl))
	c.JSON(http.StatusOK, gin.H{
		"success": res,
	})
}
