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
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"xorm.io/xorm"
)

const RESOURCES_PATH string = "./resources"
const NASU_DB_PATH string = "./resources/nasu.db"
const LOG_FILENAME string = "nasu.log"

func InitLog() {
	logFile := path.Join(RESOURCES_PATH, LOG_FILENAME)
	if res, _ := utils.IsPathOrFileExisted(logFile); !res {
		f, _ := os.Create(logFile)
		_ = f.Close()

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

func buildDb() {
	context.NasuContext.Logger.Info("[Nasu-init] Start to build sqlite3 db")
	script := exec.Command("sqlite3 " + NASU_DB_PATH)
	err := script.Run()
	if err != nil {
		context.NasuContext.Logger.Info("[Nasu-init] Fail to build sqlite3 db, please check if sqlite3 installed, err: ", err.Error())
	}
}

func BuildResourceDir() {
	context.NasuContext.Logger.Info("[Nasu-init] Start to check if resources path exists...")
	if res, err := utils.IsPathOrFileExisted(RESOURCES_PATH); res && err == nil {
		context.NasuContext.Logger.Info("[Nasu-init] Resources path exists")
		if res, err := utils.IsPathOrFileExisted(NASU_DB_PATH); !res && err == nil {
			context.NasuContext.Logger.Info("[Nasu-init] Db does not exit")
			buildDb()
		} else if res && err == nil {
			context.NasuContext.Logger.Info("[Nasu-init] Db exists")
		} else {
			context.NasuContext.Logger.Info("[Nasu-init] Fail to find db path, err: ", err.Error())
		}
	} else if !res && err == nil {
		context.NasuContext.Logger.Info("[Nasu-init] Resourcs path does not exit and try to build it...")
		err = os.Mkdir(RESOURCES_PATH, os.ModePerm)
		if err != nil {
			context.NasuContext.Logger.Info("[Nasu-init] Fail to build resources path, err: ", err.Error())
		} else {
			buildDb()
		}
	} else {
		context.NasuContext.Logger.Info("[Nasu-init] Fail to find resources path, err: ", err.Error())
	}
	absPath, _ := filepath.Abs(RESOURCES_PATH)
	context.NasuContext.ResourcesDir = absPath
}

func InitDB() (err error) {
	context.NasuContext.Logger.Info("[Nasu-init] Start to init db...")
	engine, err := xorm.NewEngine("sqlite3", NASU_DB_PATH)
	if err != nil {
		return
	}

	err = engine.Sync(new(db.NasuMeta), new(db.NasuFile))
	if err != nil {
		return
	}
	exist, _ := engine.Exist(&db.NasuMeta{
		MetaType: "PASSWORD",
	})
	if !exist {
		_, _ = engine.Insert(&db.NasuMeta{
			MetaType:  "PASSWORD",
			MetaValue: service.DEFAULT_PASSWORD, // md5 for "admin"
		})
	}
	nasuMeta := db.NasuMeta{}
	_, _ = engine.Where("meta_type = ?", "PASSWORD").Get(&nasuMeta)
	context.NasuContext.Password = nasuMeta.MetaValue
	context.NasuContext.XormEngine = engine
	context.NasuContext.Logger.Info("[Nasu-init] Db has been inited!")
	return nil
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
	var port int64
	flag.StringVar(&host, "h", "localhost", "hostname")
	flag.Int64Var(&port, "p", 8080, "port")
	gin.SetMode(gin.ReleaseMode)

	var err error
	context.NasuContext.Host = host
	context.NasuContext.Port = port
	InitLog()
	BuildResourceDir()
	err = InitDB()
	if err != nil {
		context.NasuContext.Logger.Info("[Nasu-init] Fail to init sqlite3 db, err: ", err.Error())
	}
	router := InitRoute()

	context.NasuContext.Logger.Info("[Nasu-init] Start to run App Nasu on ", context.NasuContext.Port)
	err = router.Run(":" + strconv.Itoa(int(context.NasuContext.Port)))
	if err != nil {
		context.NasuContext.Logger.Info("[Nasu-init] Fail to run App Nasu, err: ", err.Error())
		return
	}
}
