package middleware

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// LogInfo defines log structure
type LogInfo struct {
	ClientIP   *string `json:"clientIP"`
	HTTPMethod *string `json:"httpMethod"`
	URI        *string `json:"uri"`
	StatusCode *int    `json:"statusCode"`
	StartTime  *int64  `json:"startTime"`
	Cost       *int64  `json:"cost"`
}

// GetLogger returns logger
func GetLogger() gin.HandlerFunc {
	logClient := logrus.New()
	// 禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)

	logPath := "logs"
	logWriter, err := rotatelogs.New(
		logPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		start := time.Now().UnixNano() / 1e6
		c.Next()
		end := time.Now().UnixNano() / 1e6
		latency := end - start

		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logInfo := &LogInfo{}
		logInfo.ClientIP = &clientIP
		logInfo.StartTime = &start
		logInfo.Cost = &latency
		logInfo.StatusCode = &statusCode
		logInfo.URI = &path
		logInfo.HTTPMethod = &method
		logStr, err := json.Marshal(logInfo)
		if err != nil {
			panic(err)
		}
		// logClient.Infof("%v", logStr)
		logJsonStr := string(logStr)
		fmt.Println(logJsonStr)
		logClient.Infof("AccessLog: %s", logJsonStr)
		// logClient.Infof("| %3d | %13v | %15s | %s %s | %s",
		// statusCode,
		// latency,
		// clientIP,
		// method,
		// path,
		// logJsonStr,
		// )
	}
}
