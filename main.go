package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"livewallpaper/config"
	"livewallpaper/controller"
	"livewallpaper/db"
)

func init() {
	config.Setup()
	db.Setup()
}

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/videoList", controller.VideoList)
	}

	r.Run(":" + viper.GetString("port"))
}
