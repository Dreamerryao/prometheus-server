package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test400(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "400 test: bad request"})
}

func Test502(c *gin.Context) {
	c.JSON(http.StatusBadGateway, gin.H{"error": "502 test: bad gateaway"})
}

func Test500(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "500 test: internal server error"})
}

func Test200(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "200 test: OK"})
}
