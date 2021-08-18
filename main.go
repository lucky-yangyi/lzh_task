package main

import (
	"lzh-tsak/router"
	"lzh-tsak/task"
)

func main()  {
	//task.SpiderFilscan()
	go task.TaskStart()
	engine := router.Router()
	engine.Run(":8080")
}
