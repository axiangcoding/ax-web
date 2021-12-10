package service

import (
	"gin-template/internal/app/data"
	"gin-template/internal/app/data/schema"
	"gin-template/internal/app/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(ctx *gin.Context, ur entity.UserRegister) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := schema.User{
		UserName: ur.UserName,
		Email:    ur.Email,
		Phone:    ur.Phone,
		Password: string(hashedPassword),
	}
	return data.UserRegister(ctx, user)
}
