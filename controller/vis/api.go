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

func GetApis(c *gin.Context) {
	var apiCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "api")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := apiCollection.Find(ctx, bson.M{})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get api error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allApis []bson.M
	if err = result.All(ctx, &allApis); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allApis)
}
