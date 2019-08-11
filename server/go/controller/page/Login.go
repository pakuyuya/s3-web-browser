package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginGET is a function response page for "/login"
func LoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
