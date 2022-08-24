package data

import (
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/settings"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func Setup() {
	db = initDB()
}

func initDB() *gorm.DB {
	dial := selectDbDialect()
	db, err := gorm.Open(dial,
		&gorm.Config{
			NamingStrategy: &schema.NamingStrategy{SingularTable: true},
			// TODO: use project's log interface to display gorm's log
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
				logger.Config{
					SlowThreshold:             time.Second,   // 慢 SQL 阈值
					LogLevel:                  logger.Silent, // 日志级别
					IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  false,         // 禁用彩色打印
				},
			)})
	if err != nil {
		logging.Fatalf("Can't connect to Database: %s", err)
	}
	logging.Info("Database connected")
	setDbProperties(db)
	autoMigrate(db)
	return db
}

func selectDbDialect() gorm.Dialector {
	var dial gorm.Dialector
	driver := settings.Config.App.Data.Database.Driver
	source := settings.Config.App.Data.Database.Source
	// TODO: support more database driver
	switch driver {
	case "mysql":
		dial = mysql.Open(source)
		break
	case "postgres":
		dial = postgres.Open(source)
		break
	default:
		logging.Fatal("Not support such database driver yet")
	}
	return dial
}

// 自动更新表结构
func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		// 	TODO place table struct here to auto migrate
	); err != nil {
		logging.Fatal(err)
	}
	logging.Info("Auto migrate database table success")
}

func setDbProperties(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		logging.Fatal(err)
	}
	s.SetMaxOpenConns(settings.Config.App.Data.Database.MaxOpenConn)
	s.SetMaxIdleConns(settings.Config.App.Data.Database.MaxIdleConn)
}

func GetDB() *gorm.DB {
	return db
}
