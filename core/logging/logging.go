package logging

import (
	"gin-template/core/util"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Setup() *os.File {
	logPath := viper.GetString("app.filelog.path")
	if err := util.MkdirIfNotExist(logPath); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(logPath + "/all.log")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
