package data

import (
	"axiangcoding/go-gin-template/internal/app/conf"
	"axiangcoding/go-gin-template/internal/app/data/schema"
	"axiangcoding/go-gin-template/pkg/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	db = initDB()
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Config.App.Data.Database.Source),
		&gorm.Config{})
	if err != nil {
		logging.Fatal(err)
	}
	//auto migrate should not be used
	if err := db.AutoMigrate(
		&schema.User{},
	); err != nil {
		logging.Fatal(err)
	}
	logging.Info("database mysql connected success")
	s, err := db.DB()
	if err != nil {
		logging.Fatal(err)
	}
	s.SetMaxOpenConns(conf.Config.App.Data.Database.MaxOpenConn)
	s.SetMaxIdleConns(conf.Config.App.Data.Database.MaxIdleConn)

	return db
}

func GetDB() *gorm.DB {
	return db
}
