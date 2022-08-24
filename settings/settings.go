package settings

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Config AllConfig

func Setup() {
	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")
	viper.AddConfigPath("config/default/")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ax")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found. The application will not work properly")
		}
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Config Properties unable to decode into struct, %v", err)
	}
}
