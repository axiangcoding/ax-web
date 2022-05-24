package router

import (
	"github.com/axiangcoding/go-gin-template/controller/middleware"
	"github.com/axiangcoding/go-gin-template/controller/v1"
	"github.com/axiangcoding/go-gin-template/entity/app"
	"github.com/axiangcoding/go-gin-template/settings"
	"github.com/axiangcoding/go-gin-template/swagger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// r.MaxMultipartMemory = 8 << 20
	// 全局中间件
	// 使用自定义中间件
	r.Use(middleware.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	setCors(r)
	setRouterV1(r)
	return r
}

// 设置cors头
func setCors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AddAllowMethods("OPTIONS")
	config.AddAllowHeaders(app.AuthHeader)
	// r.Use(cors.New(config))
}

func setSwagger(r *gin.RouterGroup) {
	if settings.Config.App.Swagger.Enable {
		swagger.SwaggerInfo.Version = settings.Config.App.Version
		swagger.SwaggerInfo.Title = settings.Config.App.Name
		swagger.SwaggerInfo.BasePath = settings.Config.Server.BasePath
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func setRouterV1(r *gin.Engine) {
	base := r.Group(settings.Config.Server.BasePath)
	setSwagger(base)
	groupV1 := base.Group("/v1")
	{
		demo := groupV1.Group("/demo")
		{
			demo.GET("/get", v1.DemoGet)
			demo.POST("/post", v1.DemoPost)
		}
	}
}
