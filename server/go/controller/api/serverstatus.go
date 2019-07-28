package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerstatusGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "OK"})
}
