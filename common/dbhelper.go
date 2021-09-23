package common

import (
	"githhub.com/mcdexio/mai3-data/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
)

var dbLock sync.Mutex
var masterInstance *gorm.DB

func DbInstance() *gorm.DB {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()

	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}

func NewDbMaster() *gorm.DB {
	instance, err := gorm.Open("postgres", conf.Conf.DbConnStr)

	if err != nil {
		log.Fatal("open postgres error: ", err)
	}
	instance.LogMode(false)

	instance.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	masterInstance = instance

	sqlDb := masterInstance.DB()
	sqlDb.SetMaxOpenConns(0)
	sqlDb.SetMaxIdleConns(2)

	return masterInstance
}
