package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"nasu/src/context"
	"nasu/src/db"
	"nasu/src/log"
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
	log.Init(logFile)
}

func BuildResourceDir() {
	log.Log.Debug("[Nasu-init] Start to check if resources path exists...")
	if res := utils.IsPathOrFileExisted(RESOURCES_PATH); res {
		log.Log.Debug("[Nasu-init] Resources path exists")
	} else {
		log.Log.Debug("[Nasu-init] Resourcs path does not exit and try to build it...")
		err := os.Mkdir(RESOURCES_PATH, os.ModePerm)
		if err != nil {
			log.Log.Error("[Nasu-init] Fail to build resources path, err: %s", err.Error())
		}
	}
	absPath, _ := filepath.Abs(RESOURCES_PATH)
	context.NasuContext.ResourcesDir = absPath
}

func InitDB() {
	log.Log.Debug("[Nasu-init] Start to init db...")
	db.Init(NASU_DB_PATH)
	nasuMeta := db.NasuMeta{
		MetaType:  "PASSWORD",
		MetaValue: service.DEFAULT_PASSWORD, // md5 for "admin"
	}
	db.NasuMetaRepo.InsertNasuMetaIfNotExistedByMetaType(&nasuMeta)
	nasuMeta = db.NasuMeta{
		MetaType:  "HASH_PREFIX",
		MetaValue: "1",
	}
	db.NasuMetaRepo.InsertNasuMetaIfNotExistedByMetaType(&nasuMeta)
	nasuMeta = db.NasuMeta{
		MetaType:  "MAX_FILE_SIZE",
		MetaValue: "16",
	}
	db.NasuMetaRepo.InsertNasuMetaIfNotExistedByMetaType(&nasuMeta)
	context.NasuContext.Password = db.NasuMetaRepo.QueryNasuMetaByType("PASSWORD").MetaValue
	log.Log.Debug("[Nasu-init] Db has been inited!")
}

func InitRoute() *gin.Engine {
	log.Log.Debug("[Nasu-init] Start to init route...")
	router := gin.New()
	maxFileSize, _ := strconv.Atoi(db.NasuMetaRepo.QueryNasuMetaByType("MAX_FILE_SIZE").MetaValue)
	router.MaxMultipartMemory = int64(maxFileSize) << 30 // 16GB
	router.Use(gin.Recovery())
	router.Use(middleware.LogRequired())
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthRequired())
	routeAuth(authorized)
	routeCommon(router)
	log.Log.Debug("[Nasu-init] route has been inited!")
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

	log.Log.Info("[Nasu-init] Start to run App Nasu on %d", port)
	err = router.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Log.Error("[Nasu-init] Fail to run App Nasu, err: %s", err.Error())
		return
	}
}
