package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-brawser/server/go/domain/profile"
)

// ProfilesGET is a implement of WebAPI
func ProfilesGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "OK"})
}
