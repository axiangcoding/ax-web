package jwt

import (
	"gin-template/internal/app/conf"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte

type CustomerInfo struct {
	UserID string `json:"user_id"`
	Kind   string `json:"kind"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	CustomerInfo
}

func Setup() {
	jwtSecret = []byte(conf.Config.App.Token.Secret)
}

// CreateToken generate tokens used for auth
func CreateToken(username string) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
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
