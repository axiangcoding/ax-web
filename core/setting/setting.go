package setting

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() {
	setDefault()
	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")
	viper.AddConfigPath("config/default/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not foud. The program will use default setting and may not work properly")
		}
	}
	// 监控配置项，配置文件修改后会自动生效
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("Config file changed", in.Name)
	})

}

func setDefault() {
	viper.SetDefault("app.logs_path", "logs/")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.run_mode", gin.ReleaseMode)

}
