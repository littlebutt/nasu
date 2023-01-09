package log

import "github.com/sirupsen/logrus"

//go:generate mockgen -source=./log.go -destination=./log_mock.go -package=log
import (
	"fmt"
	"io"
	"os"
)

type ILog interface {
	GetLogger() *logrus.Logger
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

type log struct {
	Logger *logrus.Logger
}

var Log log

func Init(path string, isDebug bool) {
	targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Printf("已生成resources目录，你可以ctrl+C或者cmd+C结束进程后再次启动了！")
	}
	writers := []io.Writer{
		targetFile,
		os.Stderr,
	}
	textFormatter := new(logrus.TextFormatter)
	textFormatter.DisableColors = false
	textFormatter.TimestampFormat = "2006-01-02 15:04:05"
	textFormatter.FullTimestamp = true
	logger := logrus.New()
	logger.SetFormatter(textFormatter)
	logger.Out = io.MultiWriter(writers...)
	if isDebug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	Log = log{Logger: logger}
}

func (l *log) GetLogger() *logrus.Logger {
	return l.Logger
}

func (l *log) Debug(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *log) Info(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *log) Warn(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *log) Error(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *log) Fatal(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}
