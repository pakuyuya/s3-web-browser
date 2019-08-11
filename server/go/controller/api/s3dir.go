package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/profile"
	"s3-web-browser/server/go/domain/s3provider"
)

// S3dirGET is a implement as WebAPI
func S3dirGET(c *gin.Context) {
	
	profileid := c.Param("profileid")
	if profileid == "" {
		responseError(c, http.StatusBadRequest, "更新に失敗しました")
		return
	}
	path := profile.FormatBasepath(c.Param("path"))

	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	profile, err := profile.SelectByID(tx, profileid)
	if err != nil {
		panic(err)
	}

	sess, err := s3provider.CreateSession(profile.Connjson)
	if err != nil {
		panic(err)
	}
	s3items, err := s3provider.List(sess, profile.Bucket, path)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, s3items)
}
