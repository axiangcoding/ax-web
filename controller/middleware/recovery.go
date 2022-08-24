package middleware

import (
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/entity/e"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		app.ServerFailed(c, e.Error)
	})
}
