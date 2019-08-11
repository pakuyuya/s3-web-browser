package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"s3-web-browser/server/go/domain/profile"
)

// Profile is a struct for form binding
type Profile struct {
	Profilename string `form:"profilename" binding:"required"`
	Connjson string `form:"connjson" binding:"required"`
	Bucket string `form:"bucket" binding:"required"`
	Basepath string `form:"profilename"`
}

// ProfilePOST is a implement as WebAPI
func ProfilePOST(c *gin.Context) {
	form := Profile{}
	if err := c.Bind(&form); err != nil {
		responseError(c, http.StatusBadRequest, "bad request")
		return
	}

	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	p := profile.Profile {
		Profilename: form.Profilename,
		Connjson: form.Connjson,
		Bucket: form.Bucket,
		Basepath: profile.FormatBasepath(form.Basepath),
	}
	_, err = profile.Insert(tx, &p)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

// ProfilePUT is a implement as WebAPI
func ProfilePUT(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		responseError(c, http.StatusBadRequest, "更新に失敗しました")
		return
	}

	form := Profile{}
	if err := c.Bind(&form); err != nil {
		responseError(c, http.StatusBadRequest, "更新に失敗しました")
		return
	}

	conn, tx, err := getConnTx()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	p := profile.Profile {
		Profileid: id,
		Profilename: form.Profilename,
		Connjson: form.Connjson,
		Bucket: form.Bucket,
		Basepath: profile.FormatBasepath(form.Basepath),
	}
	cnt, err := profile.UpdateByID(tx, &p)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	if (cnt < 1) {
		tx.Rollback()
		responseError(c, http.StatusNotFound, "更新対象がありません")
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
