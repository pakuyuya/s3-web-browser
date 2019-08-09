package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/loginsession"
)

// LoginPOST is a implement as WebAPI
func LoginPOST(c *gin.Context) {
	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	loginid := c.PostForm("loginid")
	password := c.PostForm("password")

	logininfo, err := loginsession.Auth(tx, loginid, password)
	if err != nil {
		return responseError(c, http.StatusUnauthorized, "IDまたはパスワードが違います")
	}

	c.JSON(http.StatusOK, gin.H{
		"logininfo": logininfo,
		"result": "OK",
	})
}
