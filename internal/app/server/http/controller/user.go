package controller

import (
	"gin-template/internal/app/biz"
	"gin-template/internal/app/entity"
	"gin-template/pkg/app"
	"gin-template/pkg/app/e"
	jwt_util "gin-template/pkg/util/jwt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	uu biz.UserUseCase
}

type UserLogin struct {
	UserName    string `json:"user_name"`
	PassWordMd5 string `json:"pass_word_md_5"`
}

func (uc *UserController) UserLogin(c *gin.Context) {
	req := &UserLogin{}
	err := c.BindJSON(req)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	userLogin := &entity.UserLogin{
		UserName:    req.UserName,
		PassWordMd5: req.UserName,
	}

	userLogin, err = uc.uu.UserLogin(c, userLogin)
	if err != nil {
		app.BizFailed(c, e.ERROR, err)
		return
	}
	token, err := jwt_util.CreateToken(userLogin.UserId)
	if err != nil {
		println(err.Error())
		app.ServerFailed(c, e.ERROR, err)
		return
	}
	app.Success(c, map[string]string{
		"token": token,
	})
	return
}

type UserRegister struct {
	UserName    string `json:"user_name"`
	PassWordMd5 string `json:"pass_word_md_5"`
}

func (uc *UserController) UserRegister(c *gin.Context) {
	req := &UserLogin{}
	err := c.BindJSON(req)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	userLogin := &entity.UserLogin{
		UserName:    req.UserName,
		PassWordMd5: req.UserName,
	}

	userLogin, err = uc.uu.UserRegister(c, userLogin)
	if err != nil {
		app.BizFailed(c, e.ERROR, err)
		return
	}
	app.Success(c, map[string]string{
		"user_id": userLogin.UserId,
	})
	return
}

func NewUserController(uu biz.UserUseCase) UserController {
	return UserController{uu: uu}
}
