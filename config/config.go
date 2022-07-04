package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() {
	if gin.IsDebugging() {
		viper.SetConfigFile("./config/config_dev.yaml")
	} else {
		viper.SetConfigFile("./config/config_prod.yaml")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
