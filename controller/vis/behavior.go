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

func GetBehaviors(c *gin.Context) {
	var behaviorCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "behavior")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := behaviorCollection.Find(ctx, bson.M{})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get behavior error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allBehaviors []bson.M
	if err = result.All(ctx, &allBehaviors); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allBehaviors)
}
