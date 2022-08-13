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

func GetJsErrors(c *gin.Context) {
	var errorCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "error")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := errorCollection.Find(ctx, bson.M{"error_type": "jsError"})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get jsError error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allJsErrors []bson.M
	if err = result.All(ctx, &allJsErrors); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allJsErrors)
}

func GetResourceErrors(c *gin.Context) {
	var errorCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "error")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := errorCollection.Find(ctx, bson.M{"error_type": "resourceError"})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get resourceError error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allResourceErrors []bson.M
	if err = result.All(ctx, &allResourceErrors); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allResourceErrors)
}
