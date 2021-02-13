package dao

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

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
	db, err = gorm.Open(*conf.Db.Name, *conf.Db.URL)
	if err != nil {
		panic(err)
	}

	// Disable table name's pluralization
	db.SingularTable(*conf.Db.SingularTable)
	db.LogMode(*conf.Db.DebugMode)
	db.DB().SetMaxIdleConns(*conf.Db.MaxIdleConns)
	db.DB().SetMaxOpenConns(*conf.Db.MaxOpenConns)
}
