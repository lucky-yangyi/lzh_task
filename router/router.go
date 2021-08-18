package router

import (
	"github.com/gin-gonic/gin"
	"lzh-tsak/controller"
)

func Router()*gin.Engine{
	r := gin.New()
	r.GET("/getFilscanInfo",controller.GetFilsInfo)
	r.GET("/getPowers",controller.GetPowers)
	return r
}
