package main

import (
	"context"
	"fmt"
	"github.com/axiangcoding/go-gin-template/internal/app/conf"
	"github.com/axiangcoding/go-gin-template/internal/app/data"
	"github.com/axiangcoding/go-gin-template/pkg/auth"
	"github.com/axiangcoding/go-gin-template/pkg/cache"
	"github.com/axiangcoding/go-gin-template/pkg/logging"
	"github.com/axiangcoding/go-gin-template/pkg/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	conf.Setup()
	logging.Setup()
	data.Setup()
	cache.Setup()
	auth.Setup()
}

// @title        Golang Gin Template API
// @version      1.0.0
// @description  An example of gin
// @termsOfService

// @contact.name  axiangcoding
// @contact.url
// @contact.email  axiangcoding@gmail.com

// @license.name
// @license.url

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        token

// @accept   json
// @produce  json
func main() {
	runMode := conf.Config.Server.RunMode
	gin.SetMode(runMode)
	r := router.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Config.Server.Port),
		Handler: r,
	}

	// 在 goroutine中初始化服务器，这样就不会阻塞下文的优雅停止处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatal("Server error. ", err)
		}
	}()

	// 等待中断信号来优雅停止服务器，设置的5秒延迟
	quit := make(chan os.Signal, 1)
	// kill （不带参数的）是默认发送 syscall.SIGTERM
	// kill -2 是 syscall.SIGINT
	// kill -9 是 syscall.SIGKILL，但是无法被捕获到，所以无需添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Info("Shutting down server...")

	// ctx是用来通知服务器还有5秒的时间来结束当前正在处理的request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.Fatal("Server forced to shutdown. ", err)
	}

	logging.Info("Server exiting")
}
