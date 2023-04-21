package storage

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func open() *gorm.DB {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(path+"/rdb.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to open database")
	}
	db.AutoMigrate(&Domain{})
	db.AutoMigrate(&URL{})
	db.AutoMigrate(&DomainChange{})
	db.AutoMigrate(&UrlChange{})
	return db
}
