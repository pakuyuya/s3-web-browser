package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServerstatusGET is a function that implement as WebAPI
func ServerstatusGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "OK"})
}
