package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/profile"
)

// ProfilesGET is a implement as WebAPI
func ProfilesGET(c *gin.Context) {
	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	profiles, err := profile.SelectAll(tx)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, profiles)
}
