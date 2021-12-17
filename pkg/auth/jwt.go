package auth

import (
	"github.com/axiangcoding/go-gin-template/internal/app/conf"
	"github.com/axiangcoding/go-gin-template/internal/app/data/schema"
	"github.com/axiangcoding/go-gin-template/pkg/logging"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte
var expireDuration time.Duration

type UserInfo struct {
	UserID int64  `json:"user_id"`
	Roles  string `json:"roles"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func SetupJwt() {
	jwtSecret = []byte(conf.Config.App.Token.Secret)
	expireStr := conf.Config.App.Token.ExpireDuration
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		logging.Fatal("Config properties: app.token.expire_duration not valid")
	}
	expireDuration = expire
}

// CreateToken generate tokens used for auth
func CreateToken(user schema.User) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    conf.Config.App.Name,
		},
		UserInfo{
			UserID: user.UserId,
			Roles:  user.Roles,
		},
	}
	return t.SignedString(jwtSecret)
}

// ParseToken parsing token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
