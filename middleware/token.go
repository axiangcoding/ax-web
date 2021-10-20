package middleware

import (
	"gin-template/core/util"
	"log"
	"net/http"

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
					log.Println("That's not even a token")
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not active yet
					log.Println("Timing is everything")
				} else {
					log.Println("Couldn't handle this token:", err)
				}
			}
			log.Println("Couldn't handle this token:", err)
			c.Abort()
			return
		}
		c.Next()
		log.Println(cc.CustomerInfo)
	}
}
