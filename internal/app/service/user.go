package service

import (
	"github.com/axiangcoding/go-gin-template/internal/app/data"
	"github.com/axiangcoding/go-gin-template/internal/app/data/schema"
	"github.com/axiangcoding/go-gin-template/internal/app/entity"
	"github.com/axiangcoding/go-gin-template/pkg/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func UserRegister(ctx *gin.Context, ur entity.UserRegister) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := schema.User{
		UserName: ur.UserName,
		Email:    ur.Email,
		Phone:    ur.Phone,
		Password: string(hashedPassword),
		Roles:    schema.UserRoleNormal,
	}
	user.GenerateId()
	return data.UserRegister(ctx, user)
}

func UserLogin(ctx *gin.Context, login entity.UserLogin) (string, error) {
	user := schema.User{
		UserId: login.UserId,
	}
	findUser, err := data.UserLogin(ctx, user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(login.Password))
	if err != nil {
		return "", err
	}
	token, err := auth.CreateToken(findUser)
	if err != nil {
		return "", err
	}
	// FIXME: should use userID as key, so we can kick out expired token
	err = CacheToken(ctx, strconv.FormatInt(findUser.UserId, 10), token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func UserLogout(c *gin.Context, token string) error {
	claims, _ := auth.ParseToken(token)
	err := DeleteCachedToken(c, claims.Id)
	if err != nil {
		return err
	}
	return nil
}
