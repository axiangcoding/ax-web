package routers

import (
	v1 "gin-template/api/v1"

	docs "gin-template/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter(router *gin.Engine) {
	groupV1 := router.Group("/api/v1")
	{
		demo := groupV1.Group("/demo")
		{
			demo.GET("/get", v1.DemoGet)
			demo.POST("/post", v1.DemoPost)
		}
	}
	if viper.GetBool("swagger.enable") {
		setSwaggerInfo()
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}

func setSwaggerInfo() {
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Title = viper.GetString("app.name")
}
