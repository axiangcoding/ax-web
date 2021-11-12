// +build wireinject

package main

import (
	"gin-template/internal/app/biz"
	"gin-template/internal/app/conf"
	"gin-template/internal/app/data"
	"gin-template/internal/app/server"
	"gin-template/internal/app/server/http/controller"
	"github.com/google/wire"
)

func initApp(*conf.Data, *conf.Server) (*App, func(), error) {
	panic(
		wire.Build(
			biz.ProviderSet,
			server.ProviderSet,
			controller.ProviderSet,
			data.ProviderSet,
			newApp,
		),
	)
}
