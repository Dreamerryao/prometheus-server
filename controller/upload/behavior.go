package upload

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Dreamerryao/prometheus-server/models"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
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
	err = c.BindJSON(&pvBehavior)
	if err == nil {
		behaviorType = "pvBehavior"
	}
	err = c.BindJSON(&stayBehavior)
	if err == nil {
		behaviorType = "stayBehavior"
	}
	if behaviorType == "undefined" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "api error"})
		return
	}

	if behaviorType == "pvBehavior" {
		pvBehavior.Base.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		pvBehavior.Base.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = behaviorCollection.InsertOne(ctx, pvBehavior)
		if err != nil {
			msg := fmt.Sprintf("item not created" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if behaviorType == "stayBehavior" {
		stayBehavior.Base.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		stayBehavior.Base.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = behaviorCollection.InsertOne(ctx, stayBehavior)
		if err != nil {
			msg := fmt.Sprintf("item not created ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	c.JSON(http.StatusOK, result)
}