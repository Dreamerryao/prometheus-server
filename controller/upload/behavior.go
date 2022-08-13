package upload

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Dreamerryao/prometheus-server/models"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateBehavior(c *gin.Context) {
	var behaviorCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "behavior")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var result *mongo.InsertOneResult
	var pvBehavior models.PvBehavior
	var stayBehavior models.StayBehavior
	var behaviorType = "undefined"
	err = c.ShouldBindBodyWith(&pvBehavior, binding.JSON)
	if err == nil && pvBehavior.BehaviorType == "pv" {
		behaviorType = "pvBehavior"
	}
	err = c.ShouldBindBodyWith(&stayBehavior, binding.JSON)
	if err == nil && stayBehavior.BehaviorType == "staytime" {
		behaviorType = "stayBehavior"
	}
	if behaviorType == "undefined" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "api error"})
		return
	}

	if behaviorType == "pvBehavior" {
		pvBehavior.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		pvBehavior.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		pvBehavior.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = behaviorCollection.InsertOne(ctx, pvBehavior)
		if err != nil {
			msg := fmt.Sprintf("item not created" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if behaviorType == "stayBehavior" {
		stayBehavior.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		stayBehavior.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		stayBehavior.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = behaviorCollection.InsertOne(ctx, stayBehavior)
		if err != nil {
			msg := fmt.Sprintf("item not created ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	c.JSON(http.StatusOK, result)
}
