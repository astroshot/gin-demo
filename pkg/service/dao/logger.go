package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"gin-demo/pkg/config"
	"gin-demo/pkg/helper"
	_ "gin-demo/pkg/util"
)

var log = config.GetLogger()

var (
	infoStr       = "%s\n[info] "
	warnStr       = "%s\n[warn] "
	errStr        = "%s\n[error] "
	traceStr      = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr   = "%s %s\n[%.3fms] [rows:%v] %s"
	SlowThreshold = 300 * time.Millisecond
)

// Implement gorm Logger interface
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
// func (l *GormLogger) Print(v ...interface{}) {
// fmt.Println("v: ", v)
// vStr := fmt.Sprintf("length: %d, vals: %v", len(v), v)
// fmt.Println(vStr)
//	switch v[0] {
//	case "sql":
//		l.logger.WithFields(
//			logrus.Fields{
//				"module":       "gorm",
//				"type":         "sql",
//				"rowsReturned": v[5],
//				"src":          v[1],
//				"values":       v[4],
//				"duration":     v[2],
//			},
//		).Info(v[3])
//	case "log":
//		l.logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
//	}
//}

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

func getTraceID(ctx context.Context) string {
	return helper.GetTraceIDFrom(ctx)
}

// logger.WithFields(logrus.Fields{
// 	"proto":                c.Request.Proto,
// 	"host":                 c.Request.Host,
// 	"status":               statusCode,
// 	"method":               method,
// 	"requestContentLength": c.Request.ContentLength,
// 	"requestHeader":        c.Request.Header,
// 	"trailer":              c.Request.Trailer,
// 	"requestBody":          dataStr,
// 	"URI":                  c.Request.URL.Path,
// 	"responseBody":         blw.body.String(),
// 	"clientIP":             clientIP,
// 	"remoteAddr":           c.Request.RemoteAddr,
// 	"cost":                 latencyStr,
// 	"traceID":              traceIDStr,
// }).Infof("")
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	var traceID = getTraceID(ctx)
	l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Info(msg, data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	var traceID = getTraceID(ctx)
	l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Warn(msg, data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	var traceID = getTraceID(ctx)
	l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Error(msg, data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	var traceID = getTraceID(ctx)
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil { // && l.LogLevel >= Error:
		if rows == -1 {
			l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Printf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Printf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	} else if elapsed > SlowThreshold && SlowThreshold != 0 { // && l.LogLevel >= Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", SlowThreshold)
		if rows == -1 {
			l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Printf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Printf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	} else {
		if rows == -1 {
			l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Printf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.WithContext(ctx).WithFields(logrus.Fields{"traceID": traceID}).Printf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
	// l.logger.Tracef()
}
