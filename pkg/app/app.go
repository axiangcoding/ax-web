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

// generateErrJson generate ErrJson from errors
func generateErrJson(errors []error) ErrJson {
	hideDetail := conf.Config.App.Response.HideErrorDetails
	var errMessages []string
	if !hideDetail {
		for _, err := range errors {
			errMessages = append(errMessages, err.Error())
		}
	}
	return ErrJson{Err: errMessages}
}

// HttpResponse common response
func HttpResponse(c *gin.Context, httpCode int, msgCode int, data interface{}) {
	c.JSON(httpCode, ApiJson{
		Code: msgCode,
		Msg:  e.GetMsg(msgCode),
		Data: data,
	})
}

// Success response a success
func Success(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusOK, e.Success, data)
}

// BizFailed business failed response
func BizFailed(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 && err[0] != nil {
		HttpResponse(c, http.StatusOK, errCode, generateErrJson(err))
	} else {
		HttpResponse(c, http.StatusOK, errCode, nil)
	}
}

// BadRequest bad request response
func BadRequest(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 && err[0] != nil {
		HttpResponse(c, http.StatusBadRequest, errCode, generateErrJson(err))
	} else {
		HttpResponse(c, http.StatusBadRequest, errCode, nil)
	}
	c.Abort()
}

// ServerFailed server internal failed response
func ServerFailed(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 && err[0] != nil {
		HttpResponse(c, http.StatusInternalServerError, errCode, generateErrJson(err))
	} else {
		HttpResponse(c, http.StatusInternalServerError, errCode, nil)
	}
}

//Unauthorized authorized failed response
func Unauthorized(c *gin.Context, errCode int, err ...error) {
	if len(err) > 0 && err[0] != nil {
		HttpResponse(c, http.StatusUnauthorized, errCode, generateErrJson(err))
	} else {
		HttpResponse(c, http.StatusUnauthorized, errCode, nil)
	}
	c.Abort()
}
