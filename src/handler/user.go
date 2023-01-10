package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/context"
	"github.com/littlebutt/nasu/src/service"
	"net/http"
)

func HandleLogin(c *gin.Context) {
	password := c.PostForm("password")
	success, isFirst, token := service.Login(password)
	if success {
		context.NasuContext.Password = password
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"isFirst": isFirst,
		"token":   token,
	})
}

func HandleChangePassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	success := service.ChangePassword(oldPassword, newPassword)
	c.JSON(http.StatusOK, gin.H{
		"success": success,
	})
}
