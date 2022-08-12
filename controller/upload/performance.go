package upload

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Dreamerryao/prometheus-server/models"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePerformance(c *gin.Context) {
	var performanceCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "performance")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var result *mongo.InsertOneResult
	var timePerformance models.TimePerformance
	var paintPerformance models.PaintPerformance
	var performanceType = "undefined"
	err = c.BindJSON(&timePerformance)
	if err == nil {
		performanceType = "timePerformance"
	}
	err = c.BindJSON(&paintPerformance)
	if err == nil {
		performanceType = "paintPerformance"
	}
	if performanceType == "undefined" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "api error"})
		return
	}

	if performanceType == "timePerformance" {
		timePerformance.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		timePerformance.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		timePerformance.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = performanceCollection.InsertOne(ctx, timePerformance)
		if err != nil {
			msg := fmt.Sprintf("item not created" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if performanceType == "paintPerformance" {
		paintPerformance.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		paintPerformance.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		paintPerformance.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = performanceCollection.InsertOne(ctx, paintPerformance)
		if err != nil {
			msg := fmt.Sprintf("item not created ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	c.JSON(http.StatusOK, result)
}
