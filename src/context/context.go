package context

import (
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

type Context struct {
	Password     string
	ResourcesDir string
	XormEngine   *xorm.Engine
	Logger       *logrus.Logger
}

var NasuContext = Context{}
