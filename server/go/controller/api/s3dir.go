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
		responseError(c, http.StatusNotFound, "接続プロファイルが見つかりませんでした。")
		return
	}

	sess, err := s3provider.CreateSession(profile.Connjson)
	if err != nil {
		responseError(c, http.StatusUnauthorized, "AWSへの接続に失敗しました。接続情報を見直してください。")
		return
	}
	s3items, err := s3provider.List(sess, profile.Bucket, path)
	if err != nil {
		responseError(c, http.StatusNotFound, "S3に接続権限がないか、指定したパスが見つかりませんでした。")
		return
	}

	c.JSON(http.StatusOK, s3items)
}
