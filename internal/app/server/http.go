package server

import (
	"fmt"
	"gin-template/api/docs"
	"gin-template/internal/app/conf"
	"gin-template/internal/app/server/http/controller"
	"gin-template/internal/app/server/http/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewHTTPServer(uc controller.UserController) *http.Server {
	//初始化路由
	handler := gin.New()
	runMode := conf.Config.Server.RunMode
	gin.SetMode(runMode)
	//todo 全局中间件
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	router.RegisterRouter(handler, uc)
	if conf.Config.App.Swagger.Enable {
		setSwaggerInfo()
		handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Config.Server.Port),
		Handler: handler,
	}
	return srv
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Version = conf.Config.App.Version
	docs.SwaggerInfo.Title = conf.Config.App.Name
}
