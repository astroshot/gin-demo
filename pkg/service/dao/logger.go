package dao

import (
	_ "fmt"

	"github.com/sirupsen/logrus"

	"gin-demo/pkg/config"
)

var log = config.GetLogger()

// GormLogger struct
type GormLogger struct{}

// Print - Log Formatter
// v: [sql, invocation func position, sql execution duration, sql content, values in sql, rowsReturned]
func (*GormLogger) Print(v ...interface{}) {
	// fmt.Println("v: ", v)
	// vStr := fmt.Sprintf("length: %d, vals: %v", len(v), v)
	// fmt.Println(vStr)
	switch v[0] {
	case "sql":
		log.WithFields(
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
		log.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
