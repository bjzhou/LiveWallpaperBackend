package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"livewallpaper/config"
	"livewallpaper/db"
	"net/http"
)

func init() {
	config.Setup()
	db.Setup()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello Api V1",
			})
		})
	}

	r.Run(":" + viper.GetString("port"))
}
