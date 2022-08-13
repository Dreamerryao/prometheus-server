package main

import (
	"github.com/Dreamerryao/prometheus-server/controller/upload"
	"github.com/Dreamerryao/prometheus-server/controller/vis"
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
			uploadRoute.POST("/api", upload.CreateApi)
			uploadRoute.POST("/performance", upload.CreatePerformance)
			uploadRoute.POST("/behavior", upload.CreateBehavior)
		}
		GRoute.GET("/behavior", vis.GetBehaviors)
	}
}
