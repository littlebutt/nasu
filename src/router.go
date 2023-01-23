package main

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/handler"
	"net/http"
)

func routeCommon(engine *gin.Engine) {
	engine.POST("/login", handler.HandleLogin)

}

func routeCookie(engine *gin.RouterGroup) {
	engine.StaticFS("/upload", http.Dir("./resources"))
}

func routeAuth(engine *gin.RouterGroup) {
	engine.POST("/changePassword", handler.HandleChangePassword)
	engine.GET("/overallFileInfo", handler.HandleOverallFileInfo)
	engine.GET("/overallLabelInfo", handler.HandleOverallLabelInfo)
	engine.GET("/overallTagInfo", handler.HandleOverallTagInfo)
	engine.GET("/overallExtensionInfo", handler.HandleOverallExtensionInfo)
	engine.POST("/uploadFile", handler.HandleUploadFile)
	engine.GET("/listFilesByCondition", handler.HandleListFilesByCondition)
	engine.POST("/modifyFile", handler.HandleModifyFile)
	engine.POST("/deleteFile", handler.HandleDeleteFile)
	engine.POST("/changeHashPrefix", handler.HandleChangeHashPrefix)
	engine.POST("/changeMaxFileSize", handler.HandleChangeMaxFileSize)
	engine.POST("/changeTokenTtl", handler.HandleChangeTokenTtl)
}
