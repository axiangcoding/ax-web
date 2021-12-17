package service

import (
	"github.com/axiangcoding/go-gin-template/internal/app/data"
	"github.com/axiangcoding/go-gin-template/internal/app/data/schema"
	"github.com/axiangcoding/go-gin-template/internal/app/entity"
	jwt_util "github.com/axiangcoding/go-gin-template/pkg/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	token, err := jwt_util.CreateToken(findUser)
	if err != nil {
		return "", err
	}
	return token, nil
}
