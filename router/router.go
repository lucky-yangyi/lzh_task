package router

import (
	"github.com/gin-gonic/gin"
	"lzh-tsak/controller"
)

func Router() *gin.Engine {
	r := gin.New()
	r.GET("/api/MiningEarnings/RewardInfo", controller.GetFilsInfo)
	r.GET("/api/MiningCost", controller.GetPowers)
	return r
}
