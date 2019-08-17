package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BrowserGET is the function implemets WebAPI
func BrowserGET(c *gin.Context) {
	c.HTML(http.StatusOK, "browser.html", gin.H{})
}
