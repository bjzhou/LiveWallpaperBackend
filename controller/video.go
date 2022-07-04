package controller

import (
	"github.com/gin-gonic/gin"
	"livewallpaper/service"
	"livewallpaper/util"
	"net/http"
)

func VideoList(c *gin.Context) {
	page := util.Atoi(c.Query("page"), 1)
	pageSize := util.Atoi(c.Query("page_size"), 20)
	list, err := service.VideoList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"list": list})
	}
}
