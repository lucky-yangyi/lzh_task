package controller

import (
	"github.com/gin-gonic/gin"
	"lzh-tsak/service"
)

func GetFilsInfo(c *gin.Context) {
	ownerId := c.Query("ownerId")
	date := c.Query("date")
	resp, err := service.GetFilsInfo(ownerId, date)
	render(c, resp, err)
}

func GetPowers(c *gin.Context) {
	resp, err := service.GetListPower()
	render(c, resp, err)
}
