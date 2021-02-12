package config

import (
	"fmt"
	"os"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err: ", err)
	}

	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)

	logPath := "gin-demo"
	logWriter, err := rotatelogs.New(
		logPath+"-%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(fmt.Sprintf("%s.log", logPath)),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logger.AddHook(lfHook)
}

func GetLogger() *logrus.Logger {
	return logger
}
