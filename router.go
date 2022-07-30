package main

import (
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// GRoute总路由组
	GRoute := r.Group("/v1/api/")
	{
		// upload 埋点信息
		upload := GRoute.Group("/upload")
		{
			upload.POST("/error") //TODO: 错误上报 handler
			upload.POST("/api")
			upload.POST("/performance")
			upload.POST("/behavior")
		}
	}
}
