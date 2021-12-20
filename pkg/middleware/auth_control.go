package middleware

import (
	"github.com/axiangcoding/go-gin-template/internal/app/service"
	"github.com/axiangcoding/go-gin-template/pkg/app"
	"github.com/axiangcoding/go-gin-template/pkg/app/e"
	"github.com/axiangcoding/go-gin-template/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strconv"
	"strings"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			app.Unauthorized(c, e.TokenNotExist)
			return
		}
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					// That's not even a token
					app.Unauthorized(c, e.TokenNotValid)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// token is either expired or not active yet
					app.Unauthorized(c, e.TokenExpired)
				} else {
					app.Unauthorized(c, e.TokenNotValid)
				}
			}
			return
		}
		// check token in cache
		userID := claims.UserInfo.UserID
		cacheToken, err := service.GetCachedToken(c, strconv.FormatInt(userID, 10))
		if err != nil {
			app.Unauthorized(c, e.TokenExpired, err)
			return
		}
		if tokenString == cacheToken {
			service.RefreshTokenTTL(c, strconv.FormatInt(userID, 10))
		} else {
			app.Unauthorized(c, e.TokenExpired)
			return
		}
		// check user permission to access resource
		roles := claims.UserInfo.Roles
		roleItems := strings.Split(roles, ",")
		if len(roleItems) == 0 {
			app.Unauthorized(c, e.NoPermission)
			return
		}
		hasPermission := false
		for _, role := range roleItems {
			allowed, err := auth.GetEnforcer().Enforce(role, c.Request.URL.Path, c.Request.Method)
			if err != nil {
				app.Unauthorized(c, e.NoPermission, err)
				return
			}
			hasPermission = hasPermission || allowed
			if hasPermission {
				break
			}
		}
		if !hasPermission {
			app.Unauthorized(c, e.NoPermission)
			return
		}
		c.Next()
	}
}
