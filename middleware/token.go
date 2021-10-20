package middleware

import (
	"gin-template/core/util"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token not exist",
			})
			c.Abort()
			return
		}
		cc, err := util.ParseToken(tokenString)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					log.Warn("That's not even a token")
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					log.Warn("Timing is everything")
				} else {
					log.Warn("Couldn't handle this token:", err)
				}
			}
			log.Warn("Couldn't handle this token:", err)
			c.Abort()
			return
		}
		c.Next()
		log.Info(cc.CustomerInfo)
	}
}
