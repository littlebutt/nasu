package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/context"
	"github.com/littlebutt/nasu/src/db"
	"github.com/littlebutt/nasu/src/log"
	"github.com/littlebutt/nasu/src/middleware"
	"github.com/littlebutt/nasu/src/service"
	"github.com/littlebutt/nasu/src/utils"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

const RESOURCES_PATH string = "./resources"
const NASU_DB_PATH string = "./resources/nasu.db"
const LOG_FILENAME string = "nasu.log"

func InitLog(isDebug bool) {
	logFile := path.Join(RESOURCES_PATH, LOG_FILENAME)
	if res := utils.IsPathOrFileExisted(logFile); !res {
		f, _ := os.Create(logFile)
		defer f.Close()
	}
	log.Init(logFile, isDebug)
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
	nasuMeta = db.NasuMeta{
		MetaType:  "TOKEN_TTL",
		MetaValue: "3",
	}
	db.NasuMetaRepo.InsertNasuMetaIfNotExistedByMetaType(&nasuMeta)
	context.NasuContext.Password = db.NasuMetaRepo.QueryNasuMetaByType("PASSWORD").MetaValue
	tokenTtl, _ := strconv.Atoi(db.NasuMetaRepo.QueryNasuMetaByType("TOKEN_TTL").MetaValue)
	context.NasuContext.TokenTTL = int64(tokenTtl)
	log.Log.Debug("[Nasu-init] Db has been inited!")
}

func InitRoute(isDebug bool) *gin.Engine {
	log.Log.Debug("[Nasu-init] Start to init route...")
	router := gin.New()
	maxFileSize, _ := strconv.Atoi(db.NasuMetaRepo.QueryNasuMetaByType("MAX_FILE_SIZE").MetaValue)
	router.MaxMultipartMemory = int64(maxFileSize) << 30 // 16GB
	router.Use(gin.Recovery())
	router.Use(middleware.LogRequired())
	if isDebug {
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowHeaders = append(config.AllowHeaders, "Authorization")
		router.Use(cors.New(config))
	}
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
	var isDebug bool
	flag.StringVar(&host, "h", "localhost", "hostname")
	flag.IntVar(&port, "p", 8080, "port")
	flag.BoolVar(&isDebug, "d", false, "debug")
	flag.Parse()
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	var err error
	InitLog(isDebug)
	BuildResourceDir()
	InitDB()
	router := InitRoute(isDebug)

	log.Log.Info("[Nasu-init] Start to run App Nasu on %d", port)
	err = router.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Log.Error("[Nasu-init] Fail to run App Nasu, err: %s", err.Error())
		return
	}
}
