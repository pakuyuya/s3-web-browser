package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/db"
	"s3-web-browser/server/go/domain/profile"
)
type Profile struct {
	Profileid string `form:"profileid"`
	Profilename string `form:"profilename" binding:"required"`
	Connjson string `form:"connjson" binding:"required"`
	Bucket string `form:"bucket" binding:"required"`
	Basepath string `form:"profilename"`
}

// ProfilesPOST is a implement as WebAPI
func ProfilesPOST(c *gin.Context) {
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
