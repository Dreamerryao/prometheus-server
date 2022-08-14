package main

import (
	"github.com/Dreamerryao/prometheus-server/controller/test"
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
		GRoute.GET("/error/js", vis.GetJsErrors)
		GRoute.GET("/error/resource", vis.GetResourceErrors)
		GRoute.GET("/http", vis.GetApis)
		GRoute.POST("/performance/timing", vis.GetTimePerformanceByUrl)
		GRoute.POST("/performance/paint", vis.GetPaintPerformanceByUrl)
		testRoute := GRoute.Group("/test")
		{
			testRoute.GET("/400", test.Test400)
			testRoute.GET("/502", test.Test502)
			testRoute.GET("/500", test.Test500)
			testRoute.GET("/200", test.Test200)
		}
	}
}
