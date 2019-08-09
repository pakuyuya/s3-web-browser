package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/db"
	"s3-web-browser/server/go/domain/profile"
)

// ProfilesGET is a implement as WebAPI
func ProfilesGET(c *gin.Context) {
	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}

	profiles, err := profile.SelectAll(tx)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, profiles)
}
