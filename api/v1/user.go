package v1

import (
	"gin-template/internal/app/entity"
	"gin-template/internal/app/service"
	"gin-template/pkg/app"
	"gin-template/pkg/app/e"
	jwt_util "gin-template/pkg/util/jwt"
	"github.com/gin-gonic/gin"
)

// UserLogin
// @Summary 测试用户登录
// @Tags user
// @Param user_id query string false "user id"
// @Success 200 {string} json ""
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	userId := c.Query("user_id")
	token, err := jwt_util.CreateToken(userId)
	if err != nil {
		println(err.Error())
		c.JSON(500, gin.H{
			"token": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

type RegisterForm struct {
	UserName string
	Email    *string
	Phone    *string
	Password string
}

// UserRegister
// @Summary 用户注册
// @Tags user
// @Param form body RegisterForm false "register form"
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
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, id)
}
