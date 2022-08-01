package upload

import (
	"github.com/Dreamerryao/prometheus-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var errCollection *mongo.Collection = utils.OpenCollection(utils.ConnQuery, "error")

func CreateError(c *gin.Context) {
}
