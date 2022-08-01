package main

import (
	"github.com/Dreamerryao/prometheus-server/middlewares"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	utils.LogConfig() // 初始化日志配置
	r := gin.Default()
	r.Use(middlewares.Cors())
	initRouter(r) // 初始化路由

	err := r.Run(":8080")

	if err != nil {

		utils.Log.Fatal("服务启动失败 " + err.Error())
	}
}
