package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/db"
	"s3-web-browser/server/go/domain/profile"
)

// ProfilesGET is a implement of WebAPI
func ProfilesGET(c *gin.Context) {
	conn, err := db.Connection()
	if err != nil {
		// TODO: abort
	}
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		// TODO: abort
	}

	profiles, err := profile.SelectAll(tx)

	if err != nil {
		// TODO: abort
	}
	c.JSON(http.StatusOK, profiles)
}
