package jwt

import (
	"gin-template/core/logging"
	"gin-template/core/setting"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte
var expireDuration time.Duration

type CustomerInfo struct {
	UserID string `json:"user_id"`
	Kind   string `json:"kind"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	CustomerInfo
}

func Setup() {
	jwtSecret = []byte(setting.Config.App.Token.Secret)
	expireStr := setting.Config.App.Token.ExpireDuration
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		logging.Fatal("Config properties: app.token.expire_duration not valid")
	}
	expireDuration = expire
}

// CreateToken generate tokens used for auth
func CreateToken(username string) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "axiangcoding",
		},
		CustomerInfo{username, "human"},
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
