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

func CreateError(c *gin.Context) {
	var errCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "error")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var result *mongo.InsertOneResult
	var jsError models.JsError
	var resourceError models.ResourceError
	var errorType = "undefined"
	err = c.ShouldBindBodyWith(&jsError, binding.JSON)
	if err == nil && jsError.ErrorType == "jsError" {
		errorType = "jsError"
	}
	err = c.ShouldBindBodyWith(&resourceError, binding.JSON)
	if err == nil && resourceError.ErrorType == "resourceError" {
		errorType = "resourceError"
	}
	if errorType == "undefined" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "api error"})
		return
	}

	if errorType == "jsError" {
		jsError.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		jsError.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		jsError.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = errCollection.InsertOne(ctx, jsError)
		if err != nil {
			msg := fmt.Sprintf("item not created" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if errorType == "resourceError" {
		resourceError.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		resourceError.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		resourceError.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = errCollection.InsertOne(ctx, resourceError)
		if err != nil {
			msg := fmt.Sprintf("item not created ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	c.JSON(http.StatusOK, result)
}
