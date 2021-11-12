package user

import (
	"gin-template/internal/app/server/http/controller"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(g *gin.RouterGroup, c controller.UserController) *gin.RouterGroup {

	userG := g.Group("user")
	{
		userG.POST("login", c.Userlogin)
		userG.POST("register", c.UserRegister)
	}
	return userG
}
