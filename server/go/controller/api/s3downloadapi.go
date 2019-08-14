package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/profile"
	"s3-web-browser/server/go/domain/s3provider"
)


// S3downloadGET is a implement as WebAPI
func S3downloadGET(c *gin.Context) {
	profileid := c.Param("profileid")
	path := c.Param("path")
	if profileid == "" || path == "" {
		responseError(c, http.StatusBadRequest, "")
		return
	}

	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	profile, err := profile.SelectByID(tx, profileid)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	sess, err := s3provider.CreateSession(profile.Connjson)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	bucket := profile.Bucket
	err = s3provider.DownloadStream(sess, bucket, path, c.Writer)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
