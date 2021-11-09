package data

import (
	"gin-template/conf"
	"github.com/google/wire"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
)

// Data .
type Data struct {
	db *gorm.DB
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conf.Source), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {

	d := &Data{
		db: db,
	}
	return d, func() {

	}, nil
}
