package main

import (
	"lzh-tsak/config"
	"lzh-tsak/router"
	"lzh-tsak/task"
)

func main() {
	task.SpiderFilscan()
	go task.TaskStart()
	engine := router.Router()
	engine.Run(config.Config.Port)
}
