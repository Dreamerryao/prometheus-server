package main

import (
	"sync"

	"github.com/Dreamerryao/prometheus-server/middlewares"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	configInit(wg)

	r := gin.Default()
	r.Use(middlewares.Cors())
	initRouter(r) // 初始化路由
	err := r.Run(":8080")
	if err != nil {
		wg.Wait()
		utils.Log.Fatal("服务启动失败 " + err.Error())
	}
}

func configInit(wg *sync.WaitGroup) {
	utils.LogConfig() // 初始化日志配置
	go func() {
		utils.MongoDBInit() // mongodb 初始化，连接数据库
		wg.Done()
	}()
}
