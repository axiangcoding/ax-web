package v1

import (
	"gin-template/internal/app/server/http/controller"
	"gin-template/internal/app/server/http/router/v1/user"
	"github.com/gin-gonic/gin"
)

func NewV1Router(g *gin.Engine,
	uc controller.UserController,
) *gin.RouterGroup {
	v1G := g.Group("api/v1")
	user.NewUserRouter(v1G, uc)
	return v1G
}
