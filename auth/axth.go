package auth

import (
	"github.com/axiangcoding/axth"
	"github.com/axiangcoding/go-gin-template/logging"
	"github.com/axiangcoding/go-gin-template/settings"
)

var axthEnforcer *axth.Enforcer

func SetupAxth() {
	e, err := axth.NewEnforcer(&axth.Config{DBDsn: settings.Config.Data.Database.Source})
	if err != nil {
		logging.Fatal(err)
	}
	axthEnforcer = e
}

func GetAxthEnforcer() *axth.Enforcer {
	return axthEnforcer
}
