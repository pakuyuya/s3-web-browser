package api

import (
	"net/http"

	"s3-web-browser/server/go/domain/profile"

	"github.com/gin-gonic/gin"
)

// ProfileResoponse is struct for response
type ProfileResoponse struct {
	Profileid   string `json:"profileid"`
	Profilename string `json:"profilename"`
}

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

	responses := make([]ProfileResoponse, 0)
	for _, profile := range profiles {
		r := ProfileResoponse{Profileid: profile.Profileid, Profilename: profile.Profilename}
		responses = append(responses, r)
	}
	tx.Rollback()
	c.JSON(http.StatusOK, responses)
}
