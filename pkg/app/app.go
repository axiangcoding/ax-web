package app

import (
	"gin-template/pkg/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ErrJson struct {
	Err string `json:"err"`
}

// HttpResponse  setting gin.JSON
func HttpResponse(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, ApiJson{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// BizFailed 业务错误
func BizFailed(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 {
		HttpResponse(c, http.StatusOK, errCode, ErrJson{Err: err[0].Error()})
	} else {
		HttpResponse(c, http.StatusOK, errCode, nil)
	}
}

//ServerFailed 服务内部错误
func ServerFailed(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 {
		HttpResponse(c, http.StatusInternalServerError, errCode, ErrJson{Err: err[0].Error()})
	} else {
		HttpResponse(c, http.StatusInternalServerError, errCode, nil)
	}
}

//Unauthorized 权限验证
func Unauthorized(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusUnauthorized, http.StatusUnauthorized, data)
	c.Abort()
	return
}

// Success 成功返回
func Success(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusOK, e.SUCCESS, data)
	return
}

// CasbinAuthFail casbin 鉴权失败，返回 405 方法不允许访问
func CasbinAuthFail(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, data)
	c.Abort()
}

// BadRequest 请求失败
func BadRequest(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 {
		HttpResponse(c, http.StatusBadRequest, errCode, &ErrJson{Err: err[0].Error()})
		c.Abort()
	} else {
		HttpResponse(c, http.StatusBadRequest, errCode, nil)
		c.Abort()
	}
}
