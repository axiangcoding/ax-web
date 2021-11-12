package data

import (
	"gin-template/internal/app/conf"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewUserRepo,
)

// Data .
type Data struct {
	db *gorm.DB
}

func NewDB(conf *conf.Data) *gorm.DB {

	db, err := gorm.Open(mysql.Open(conf.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB) (*Data, func(), error) {
	d := &Data{
		db: db,
	}
	return d, func() {

	}, nil
}
