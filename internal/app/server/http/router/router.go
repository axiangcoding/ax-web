package router

import (
	"gin-template/internal/app/server/http/controller"
	v1 "gin-template/internal/app/server/http/router/v1"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, uc controller.UserController) {
	//sayHai
	r.HEAD("/", sayHai)
	r.GET("/", sayHai)
	pprof.Register(r)
	v1.NewV1Router(r, uc)
}
func sayHai(c *gin.Context) {
	_, err := c.Writer.Write([]byte("hello!"))
	if err != nil {
		return
	}
	return
}
