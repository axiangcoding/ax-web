package controller

import (
	"gin-template/internal/app/biz"
	jwt_util "gin-template/pkg/util/jwt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uu biz.UserUseCase
}

func (uc *UserController) Userlogin(c *gin.Context) {
	user_id := c.Query("user_id")
	token, err := jwt_util.CreateToken(user_id)
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

func (uc *UserController) UserRegister(c *gin.Context) {

}

func NewUserController(uu biz.UserUseCase) UserController {
	return UserController{uu: uu}
}
