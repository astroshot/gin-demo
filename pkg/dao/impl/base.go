package impl

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"astroshot/gin-demo/pkg/config"
)

var (
	db   *gorm.DB
	once sync.Once
)

type BaseDAOImpl struct {
	db *gorm.DB
}

func GetDB() *gorm.DB {
	return db
}

func initDB() {
	var err error

	conf := config.GetConfig(config.GetEnv())
	db, err = gorm.Open(*conf.Db.Name, *conf.Db.Url)

	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
}

func init() {
	once.Do(initDB())
}
