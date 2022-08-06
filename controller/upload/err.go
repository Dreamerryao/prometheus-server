package upload

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Dreamerryao/prometheus-server/common"
	"github.com/Dreamerryao/prometheus-server/models"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
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
	err = c.BindJSON(&jsError)
	if err == nil {
		errorType = "jsError"
	}
	err = c.BindJSON(&resourceError)
	if err == nil {
		errorType = "resourceError"
	}
	if errorType == "undefined" {
		c.JSON(http.StatusOK, common.Response{})
		return
	}

	if errorType == "jsError" {
		jsError.Base.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		jsError.Base.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = errCollection.InsertOne(ctx, jsError)
		if err != nil {
			msg := fmt.Sprintf("item not created" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if errorType == "resourceError" {
		resourceError.Base.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		resourceError.Base.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, err = errCollection.InsertOne(ctx, resourceError)
		if err != nil {
			msg := fmt.Sprintf("item not created ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	c.JSON(http.StatusOK, result)
}
