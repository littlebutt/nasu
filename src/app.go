package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"io"
	"nasu/src/context"
	"nasu/src/db"
	"nasu/src/middleware"
	"nasu/src/service"
	"nasu/src/utils"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

const RESOURCES_PATH string = "./resources"
const NASU_DB_PATH string = "./resources/nasu.db"
const LOG_FILENAME string = "nasu.log"

func InitLog() {
	logFile := path.Join(RESOURCES_PATH, LOG_FILENAME)
	if res := utils.IsPathOrFileExisted(logFile); !res {
		f, _ := os.Create(logFile)
		defer f.Close()
	}
	targetFile, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Printf("[Nasu-Log] Fail to open log file, filename: %s, err: %s\n", logFile, err.Error())
	}
	writers := []io.Writer{
		targetFile,
		os.Stderr,
	}
	logger := logrus.New()
	logger.Out = io.MultiWriter(writers...)
	logger.SetLevel(logrus.DebugLevel)
	context.NasuContext.Logger = logger
}

func BuildResourceDir() {
	context.NasuContext.Logger.Info("[Nasu-init] Start to check if resources path exists...")
	if res := utils.IsPathOrFileExisted(RESOURCES_PATH); res {
		context.NasuContext.Logger.Info("[Nasu-init] Resources path exists")
	} else {
		context.NasuContext.Logger.Info("[Nasu-init] Resourcs path does not exit and try to build it...")
		err := os.Mkdir(RESOURCES_PATH, os.ModePerm)
		if err != nil {
			context.NasuContext.Logger.Info("[Nasu-init] Fail to build resources path, err: ", err.Error())
		}
	}
	absPath, _ := filepath.Abs(RESOURCES_PATH)
	context.NasuContext.ResourcesDir = absPath
}

func InitDB() {
	context.NasuContext.Logger.Info("[Nasu-init] Start to init db...")
	engine := db.NewEngine(NASU_DB_PATH)
	context.NasuContext.XormEngine = engine
	nasuMeta := db.NasuMeta{
		MetaType:  "PASSWORD",
		MetaValue: service.DEFAULT_PASSWORD, // md5 for "admin"
	}
	db.InsertNasuMetaIfNotExistedByMetaType(&nasuMeta)
	context.NasuContext.Password = db.QueryNasuMetaByType("PASSWORD").MetaValue
	context.NasuContext.Logger.Info("[Nasu-init] Db has been inited!")
}

func InitRoute() *gin.Engine {
	context.NasuContext.Logger.Info("[Nasu-init] Start to init route...")
	router := gin.New()
	// TODO: customize maximum uploading file size
	router.MaxMultipartMemory = 16 << 30 // 16GB
	router.Use(gin.Recovery())
	router.Use(middleware.LogRequired())
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthRequired())
	routeAuth(authorized)
	routeCommon(router)
	context.NasuContext.Logger.Info("[Nasu-init] route has been inited!")
	return router
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "h", "localhost", "hostname")
	flag.IntVar(&port, "p", 8080, "port")
	gin.SetMode(gin.ReleaseMode)

	var err error
	InitLog()
	BuildResourceDir()
	InitDB()
	router := InitRoute()

	context.NasuContext.Logger.Info("[Nasu-init] Start to run App Nasu on ", port)
	err = router.Run(":" + strconv.Itoa(port))
	if err != nil {
		context.NasuContext.Logger.Info("[Nasu-init] Fail to run App Nasu, err: ", err.Error())
		return
	}
}
