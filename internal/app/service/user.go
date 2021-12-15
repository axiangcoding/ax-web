package service

import (
	"axiangcoding/go-gin-template/internal/app/data"
	"axiangcoding/go-gin-template/internal/app/data/schema"
	"axiangcoding/go-gin-template/internal/app/entity"
	jwt_util "axiangcoding/go-gin-template/pkg/util/jwt"
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
	token, err := jwt_util.CreateToken(strconv.FormatInt(findUser.UserId, 10))
	if err != nil {
		return "", err
	}
	return token, nil
}
