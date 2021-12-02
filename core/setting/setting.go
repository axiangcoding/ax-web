package setting

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Config AllConfig

func Setup() {
	setDefault()
	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")
	viper.AddConfigPath("config/default/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found. The program will use default setting and may not work properly")
		}
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Config Properties unable to decode into struct, %v", err)
	}

}

func setDefault() {
	// generate a default config file maybe?
	// 是否要生成一个默认的配置文件？
	viper.SetDefault("app.version", "0.0.1")
	viper.SetDefault("app.name", "axiangcoding/go-gin-template")
	viper.SetDefault("app.log.level", "INFO")
	viper.SetDefault("app.log.file.enable", "false")
	viper.SetDefault("app.log.file.path", "./logs/")
	viper.SetDefault("app.token,secret", "randomSecret")
	viper.SetDefault("app.token.expire_duration", "1h")
	viper.SetDefault("app.swagger.enable", true)
	viper.SetDefault("server.run_mode", gin.ReleaseMode)
	viper.SetDefault("server.port", 8080)
}
