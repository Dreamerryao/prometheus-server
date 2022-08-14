package vis

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UrlStruct struct {
	Url string `json:"url" binding:"required"`
}

func GetTimePerformanceByUrl(c *gin.Context) {
	var performanceCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "performance")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url UrlStruct
	err := c.ShouldBind(&url)
	result, err := performanceCollection.Find(ctx, bson.M{"url": url.Url, "perf_type": "timing"})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get time performance by url error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allPerformances []bson.M
	if err = result.All(ctx, &allPerformances); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allPerformances)
}

func GetPaintPerformanceByUrl(c *gin.Context) {
	var performanceCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "performance")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var url UrlStruct
	err := c.ShouldBind(&url)
	result, err := performanceCollection.Find(ctx, bson.M{"url": url.Url, "perf_type": "paint"})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get time performance by url error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allPerformances []bson.M
	if err = result.All(ctx, &allPerformances); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allPerformances)
}
