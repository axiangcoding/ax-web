package util

import "github.com/spf13/viper"

func Setup() {
	jwtSecret = []byte(viper.GetString("app.token.secret"))
}
