package auth

import (
	"github.com/axiangcoding/go-gin-template/data"
	"github.com/axiangcoding/go-gin-template/logging"
	"github.com/axiangcoding/go-gin-template/settings"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"path"
)

var casbinEnforcer *casbin.Enforcer

func SetupCasbin() {

	modelPath := settings.Config.Auth.CasbinModelPath
	policyPath := settings.Config.Auth.CasbinPolicyPath
	adapter := settings.Config.Auth.CasbinPolicyAdapter
	var e *casbin.Enforcer

	switch adapter {
	case "db":
		// 使用 gormadapter 来从数据库中加载policy，并且使用的是已经存在的gorm实例
		adapter, err := gormadapter.NewAdapterByDB(data.GetDB())
		if err != nil {
			logging.Fatal(err)
		}
		e, err = casbin.NewEnforcer("config/default/model.conf", adapter)
		if err != nil {
			logging.Fatal(err)
		}
		break
	case "local":
		var err error
		e, err = casbin.NewEnforcer(path.Join(modelPath, "model.conf"), path.Join(policyPath, "/policy.csv"))
		if err != nil {
			logging.Fatal(err)
		}
		break
	default:
		logging.Fatal("Wrong config to casbin policy")
		break
	}
	err := e.LoadPolicy()
	if err != nil {
		logging.Fatal(err)
	}
	casbinEnforcer = e
}

func GetCasbinEnforcer() *casbin.Enforcer {
	return casbinEnforcer
}
