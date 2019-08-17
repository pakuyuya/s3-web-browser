package api

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"s3-web-browser/server/go/domain/loginsession"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginPOST is a implement as WebAPI
func LoginPOST(c *gin.Context) {
	session := sessions.Default(c)

	var form struct {
		Loginid  string `form:"loginid" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	if err := c.Bind(&form); err != nil {
		responseError(c, http.StatusBadRequest, "bad request")
		return
	}

	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	logininfo, err := loginsession.Auth(tx, form.Loginid, form.Password)
	if err != nil {
		fmt.Println(err.Error())
		responseError(c, http.StatusUnauthorized, "IDまたはパスワードが違います")
		tx.Rollback()
		return
	}

	gob.Register(logininfo)
	session.Set(loginsession.SessionKey, logininfo)
	err = session.Save()
	if err != nil {
		panic(err)
	}

	tx.Rollback()
	c.JSON(http.StatusOK, gin.H{
		"redirectTo": "browser",
		"logininfo":  logininfo,
		"result":     "OK",
	})
}
