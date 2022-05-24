package auth

import (
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/axiangcoding/axth"
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
