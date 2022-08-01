package main

import (
	"github.com/Dreamerryao/prometheus-server/controller/upload"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// GRoute总路由组
	GRoute := r.Group("/v1/api/")
	{
		// upload 埋点信息
		uploadRoute := GRoute.Group("/upload")
		{
			uploadRoute.POST("/error", upload.CreateError) //TODO: 错误上报 handler
			uploadRoute.POST("/api")
			uploadRoute.POST("/performance")
			uploadRoute.POST("/behavior")
		}
	}
}
