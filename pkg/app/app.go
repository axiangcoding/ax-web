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
	Err string `json:"err"`
}

// Success response a success
func Success(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusOK, e.Success, data)
}

// HttpResponse  common response
func HttpResponse(c *gin.Context, httpCode int, msgCode int, data interface{}) {
	c.JSON(httpCode, ApiJson{
		Code: msgCode,
		Msg:  e.GetMsg(msgCode),
		Data: data,
	})
}

// BizFailed business failed response
func BizFailed(c *gin.Context, errCode int, err ...error) {
	hideDetail := conf.Config.App.Response.HideErrorDetails
	if len(err) > 0 && err[0] != nil && !hideDetail {
		HttpResponse(c, http.StatusOK, errCode, ErrJson{Err: err[0].Error()})
	} else {
		HttpResponse(c, http.StatusOK, errCode, nil)
	}
}

// ServerFailed server internal failed response
func ServerFailed(c *gin.Context, errCode int, err ...error) {
	hideDetail := conf.Config.App.Response.HideErrorDetails
	if len(err) > 0 && err[0] != nil && !hideDetail {
		HttpResponse(c, http.StatusInternalServerError, errCode, ErrJson{Err: err[0].Error()})
	} else {
		HttpResponse(c, http.StatusInternalServerError, errCode, nil)
	}
}

//Unauthorized authorized failed response
func Unauthorized(c *gin.Context, errCode int, err ...error) {
	hideDetail := conf.Config.App.Response.HideErrorDetails
	if len(err) > 0 && err[0] != nil && !hideDetail {
		HttpResponse(c, http.StatusUnauthorized, errCode, ErrJson{Err: err[0].Error()})
	} else {
		HttpResponse(c, http.StatusUnauthorized, errCode, nil)
	}
	c.Abort()
}
