package dao

import (
	"context"
	_ "fmt"
	"time"

	"github.com/sirupsen/logrus"
	glogger "gorm.io/gorm/logger"

	"gin-demo/pkg/config"
)

var log = config.GetLogger()

// TODO: implement this interface
// type Interface interface {
// 	LogMode(LogLevel) Interface
// 	Info(context.Context, string, ...interface{})
// 	Warn(context.Context, string, ...interface{})
// 	Error(context.Context, string, ...interface{})
// 	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
// }

// GormLogger struct
type GormLogger struct {
	logger *logrus.Logger
}

// GetGormLogger returns logger for gorm
func GetGormLogger() *GormLogger {
	return &GormLogger{
		logger: config.GetLogger(),
	}
}

// Print - Log Formatter
// v: [sql, invocation func position, sql execution duration, sql content, values in sql, rowsReturned]
func (l *GormLogger) Print(v ...interface{}) {
	// fmt.Println("v: ", v)
	// vStr := fmt.Sprintf("length: %d, vals: %v", len(v), v)
	// fmt.Println(vStr)
	switch v[0] {
	case "sql":
		l.logger.WithFields(
			logrus.Fields{
				"module":       "gorm",
				"type":         "sql",
				"rowsReturned": v[5],
				"src":          v[1],
				"values":       v[4],
				"duration":     v[2],
			},
		).Info(v[3])
	case "log":
		l.logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

// LogMode implements gorm Logger Interface
func (l *GormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	newLogger := l
	if level == glogger.Silent {
		newLogger.logger.Level = logrus.PanicLevel
	} else if level == glogger.Error {
		newLogger.logger.Level = logrus.ErrorLevel
	} else if level == glogger.Warn {
		newLogger.logger.Level = logrus.WarnLevel
	} else if level == glogger.Info {
		newLogger.logger.Level = logrus.InfoLevel
	} else {
		newLogger.logger.Level = logrus.DebugLevel
	}

	return newLogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Infof(msg, data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warnf(msg, data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Errorf(msg, data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	// l.logger.Tracef()
}
