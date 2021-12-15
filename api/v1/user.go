package v1

import (
	"github.com/axiangcoding/go-gin-template/internal/app/entity"
	"github.com/axiangcoding/go-gin-template/internal/app/service"
	"github.com/axiangcoding/go-gin-template/pkg/app"
	"github.com/axiangcoding/go-gin-template/pkg/app/e"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	UserId   int64
	Password string
}

// UserLogin
// @Summary User login
// @Tags user
// @Param form body LoginForm true "register form"
// @Success 200 {object} app.ApiJson ""
// @Failure 500 {object} app.ErrJson ""
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	form := LoginForm{}
	err := c.BindJSON(&form)
	if err != nil {
		app.BizFailed(c, e.RequestParamsNotValid, err)
		return
	}

	login := entity.UserLogin{
		UserId:   form.UserId,
		Password: form.Password,
	}

	token, err := service.UserLogin(c, login)
	if err != nil {
		app.BizFailed(c, e.LoginFailed, err)
		return
	}
	app.Success(c, map[string]string{"token": token})
}

type RegisterForm struct {
	UserName string
	Email    *string
	Phone    *string
	Password string
}

// UserRegister
// @Summary User register
// @Tags user
// @Param form body RegisterForm true "register form"
// @Success 200 {object} app.ApiJson ""
// @Failure 500 {object} app.ErrJson ""
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	regForm := RegisterForm{}
	err := c.BindJSON(&regForm)
	if err != nil {
		app.BizFailed(c, e.RequestParamsNotValid, err)
		return
	}
	register := entity.UserRegister{
		UserName: regForm.UserName,
		Email:    regForm.Email,
		Phone:    regForm.Phone,
		Password: regForm.Password,
	}
	id, err := service.UserRegister(c, register)
	if err != nil {
		app.BizFailed(c, e.RegisterFailed, err)
		return
	}
	app.Success(c, map[string]string{"id": id})
}
