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

func CreateApi(c *gin.Context) {
	var apiCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "api")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var result *mongo.InsertOneResult
	var api models.Api
	err = c.ShouldBindBodyWith(&api, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "api error"})
		return
	}
	api.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	api.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	api.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result, err = apiCollection.InsertOne(ctx, api)
	if err != nil {
		msg := fmt.Sprintf("item not created" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, result)
}
