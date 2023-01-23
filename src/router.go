package main

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/handler"
	"io/fs"
	"net/http"
)

func getAssets(engine *gin.Engine, fs fs.FS, filepath string) {
	engine.GET(filepath, func(context *gin.Context) {
		context.FileFromFS(filepath, http.FS(fs))
	})
}

func routeCommon(engine *gin.Engine, fileSystem fs.FS) {
	engine.GET("/", func(context *gin.Context) {
		context.FileFromFS("/index.htm", http.FS(fileSystem))
	})
	// FIXME: more elegant way to solve path mapping
	getAssets(engine, fileSystem, "/asset-manifest.json")
	getAssets(engine, fileSystem, "/favicon.ico")
	getAssets(engine, fileSystem, "/logo.svg")
	getAssets(engine, fileSystem, "/manifest.json")
	f, _ := fs.Sub(fileSystem, "static")
	engine.StaticFS("/static", http.FS(f))
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
