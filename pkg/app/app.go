package app

import (
	"github.com/axiangcoding/go-gin-template/internal/app/conf"
	"github.com/axiangcoding/go-gin-template/pkg/app/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrJson struct {
	Err []string `json:"err"`
}

// generateErrJson
// generate ErrJson from errors
// 从errors中生成 ErrJson
func generateErrJson(errors []error) *ErrJson {
	if len(errors) == 0 {
		return nil
	}

	hideDetail := conf.Config.App.Response.HideErrorDetails
	var errMessages []string
	if !hideDetail {
		for _, err := range errors {
			errMessages = append(errMessages, err.Error())
		}
	}
	return &ErrJson{Err: errMessages}
}

// HttpResponse
// common response
// 返回通用
func HttpResponse(c *gin.Context, httpCode int, msgCode int, data interface{}) {
	c.JSON(httpCode, ApiJson{
		Code: msgCode,
		Msg:  e.CodeText(msgCode),
		Data: data,
	})
}

// Success
// response a success
// 返回成功
func Success(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusOK, e.Success, data)
}

// BizFailed
// business failed response
// 返回业务逻辑失败
func BizFailed(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusOK, errCode, generateErrJson(err))
}

// BadRequest
// bad request response
// 返回错误参数请求
func BadRequest(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusBadRequest, errCode, generateErrJson(err))
	c.Abort()
}

// ServerFailed
// server internal failed response
// 返回服务器内部故障
func ServerFailed(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusInternalServerError, errCode, generateErrJson(err))
}

// Unauthorized
// authorized failed response
// 返回权限不足
func Unauthorized(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusUnauthorized, errCode, generateErrJson(err))
	c.Abort()
}

// Forbidden
// authorized forbidden response
// 返回被禁止访问
func Forbidden(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusForbidden, errCode, generateErrJson(err))
	c.Abort()
}
