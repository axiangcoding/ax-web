# Go Gin Template [![Go Report Card](https://goreportcard.com/badge/gitee.com/axiangcoding/go-gin-template)](https://goreportcard.com/report/gitee.com/axiangcoding/go-gin-template) [![MIT](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE)

切换语言: 中文|[English](./README.md)

## 介绍

使用 [Gin](https://github.com/gin-gonic/gin) 构建web应用的一个初始模板，基本涉及到了日常开发中所涉及到的技术和思路，定义了项目的结构，并提供了一些简易的样例以供使用。当然我们要认识到，Go是一个较为开放的语言，Gin也提供了十分简洁的API以供使用，因此项目的结构也可以根据实际情况进行调整，规范化的结构只是为了团队能够按照统一的风格编码，提供最佳的编码体验


## 如何运行

### 环境

1. 首先，请确保你的机器上已经安装了Go，安装的方式可以参考官网 https://golang.org/dl/

2. [TODO] 同时，由于本项目还采用了Docker来建立开发环境，如果需要确保Demo运行正常，请确保你的机器上已经安装了Docker。安装的方式可以参考官网 https://www.docker.com/get-started

3. 当然，如果你十分清楚你需要些什么服务，也可以使用自建服务（如Redis，Mysql等）来代替Docker提供的测试环境

### 配置
本项目对配置文件的读取遵循以下规则：
- 如果`config/app.toml`配置文件存在，则读取`config/app.toml`作为配置文件
- 如果不存在，则使用默认配置文件`config/default/app.toml`

因此，建议您将`config/default/app.toml`复制到`config/app.toml`，并修改`config/default/app.toml`，这样在错误修改了配置文件后，也可以快速复原。配置文件内容大致如下：

```
[app]
version = "0.0.1"
name = "Gin开发基础模板"

[app.filelog]
enable = false
path = "./logs/"

[server]
run_mode = "debug"
port = 8080

···
```

### 运行
```
$ git clone https://gitee.com/axiangcoding/go-gin-template.git
$ cd go-gin-template
$ # [TODO] automatic start docker-compose
$ go run main.go
```

项目的运行信息将会出现在你的控制台中：

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/demo/get          --> gin-template/api/v1.DemoGet (3 handlers)
[GIN-debug] POST   /api/v1/demo/post         --> gin-template/api/v1.DemoPost (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
```
接着，访问地址：[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)，你将看到如下界面：
![image](docs/swagger-image.png)

接下来，就可以使用啦

## 特性

- Gin
- Swagger
- Viper
- Graceful stop
- App configurable

## 参考

站在巨人的肩膀上，才能看得更远

- https://github.com/eddycjy/go-gin-example/blob/master/README.md
- https://sodocumentation.net/go/topic/9463/best-practices-on-project-structure
- https://docs.docker.com/get-docker/
- https://gin-gonic.com/docs/
- https://github.com/swaggo/swag/blob/master/README.md
- https://github.com/spf13/viper/blob/master/README.md
- https://pkg.go.dev/github.com/golang-jwt/jwt#section-documentation

## 联系我

邮件发送至<wyxworkmail@163.com>

