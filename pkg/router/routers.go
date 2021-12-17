package router

import (
	"github.com/axiangcoding/go-gin-template/api/docs"
	v1 "github.com/axiangcoding/go-gin-template/api/v1"
	"github.com/axiangcoding/go-gin-template/internal/app/conf"
	"github.com/axiangcoding/go-gin-template/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 全局中间件
	// 使用自定义中间件
	r.Use(middleware.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	groupV1 := r.Group("/api/v1")
	{
		demo := groupV1.Group("/demo", middleware.Token())
		{
			demo.GET("/get", v1.DemoGet)
			demo.POST("/post", v1.DemoPost)
		}
		test := groupV1.Group("/test")
		{
			test.GET("/test-log", v1.TestLog)
		}
		user := groupV1.Group("/user")
		{
			user.POST("/login", v1.UserLogin)
			user.POST("/register", v1.UserRegister)
		}

	}

	if conf.Config.App.Swagger.Enable {
		setSwaggerInfo()
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return r
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Version = conf.Config.App.Version
	docs.SwaggerInfo.Title = conf.Config.App.Name
	docs.SwaggerInfo.BasePath = conf.Config.Server.BasePath
}
