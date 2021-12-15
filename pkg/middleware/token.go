package middleware

import (
	"gin-template/pkg/app"
	"gin-template/pkg/app/e"
	jwt_util "gin-template/pkg/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			app.Unauthorized(c, e.TokenNotExist)
			return
		}
		_, err := jwt_util.ParseToken(tokenString)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					// That's not even a token
					app.Unauthorized(c, e.TokenNotLegal)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					app.Unauthorized(c, e.TokenExpired)
				} else {
					app.Unauthorized(c, e.TokenNotLegal)
				}
			}
			return
		}
		c.Next()
	}
}
