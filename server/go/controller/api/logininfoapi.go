package api

import (
	"net/http"

	"s3-web-browser/server/go/domain/loginsession"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LogininfoGET is a implement as WebAPI
func LogininfoGET(c *gin.Context) {
	session := sessions.Default(c)

	v := session.Get(loginsession.SessionKey)
	if v == nil {
		responseError(c, http.StatusNotFound, "")
		return
	}

	c.JSON(http.StatusOK, v)
}
