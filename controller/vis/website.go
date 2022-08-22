package vis

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Dreamerryao/prometheus-server/models"
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func CreateWebsite(c *gin.Context) {
	var webCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "website")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var result *mongo.InsertOneResult
	var website models.Website
	err = c.ShouldBind(&website)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "api error"})
		return
	}
	website.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	result, err = webCollection.InsertOne(ctx, website)
	if err != nil {
		msg := fmt.Sprintf("item not created" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetWebsites(c *gin.Context) {
	var webCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "website")
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := webCollection.Find(ctx, bson.M{})
	if err != nil && err.Error() != "document is nil" {
		msg := fmt.Sprintf("get api error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var allWebsites []bson.M
	if err = result.All(ctx, &allWebsites); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allWebsites)
}
