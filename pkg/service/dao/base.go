package dao

import (
	// "log"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin-demo/pkg/config"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	return db
}

func initDB() {
	var err error

	conf := config.GetConfig(config.GetEnv())

	db, err = gorm.Open(mysql.Open(*conf.Db.URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Disable table name's pluralization
	// db.SingularTable(*conf.Db.SingularTable)
	// db.LogMode(*conf.Db.DebugMode)
	dbLogger := &GormLogger{}

	// db.SetLogger(dbLogger)
	// Formatter := new(dbLogger.TextFormatter)
	// log.SetFormatter(Formatter)
	// Formatter.FullTimestamp = true
	// db.SetLogger(gorm.Logger{revel.TRACE})
	// db.DB().SetMaxIdleConns(*conf.Db.MaxIdleConns)
	// db.DB().SetMaxOpenConns(*conf.Db.MaxOpenConns)
}
