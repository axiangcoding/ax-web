package auth

import (
	"github.com/axiangcoding/go-gin-template/pkg/logging"
	"github.com/casbin/casbin/v2"
)

var enforcer *casbin.Enforcer

func SetupCasbin() {
	e, err := casbin.NewEnforcer("config/default/model.conf", "config/default/policy.csv")
	if err != nil {
		logging.Error(err)
	}
	enforcer = e
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
