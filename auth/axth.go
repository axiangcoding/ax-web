package auth

import (
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/axiangcoding/axth"
)

var axthEnforcer *axth.Enforcer

func SetupAxth() {
	options, err := axth.DefaultOptions(settings.Config.Data.Database.Source)
	if err != nil {
		logging.Fatal(err)
	}
	e, err := axth.NewEnforcer(options)
	if err != nil {
		logging.Fatal(err)
	}
	axthEnforcer = e
}

func GetAxthEnforcer() *axth.Enforcer {
	return axthEnforcer
}
