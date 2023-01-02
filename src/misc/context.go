package misc

import (
	"github.com/sirupsen/logrus"
	"sync"
	"xorm.io/xorm"
)

type Context struct {
	Password     string
	ResourcesDir string
	Host         string
	Port         int64
	XormEngine   *xorm.Engine
	Logger       *logrus.Logger
}

var instance *Context

var once sync.Once

func GetContextInstance() *Context {
	once.Do(func() {
		instance = &Context{}
	})
	return instance
}
