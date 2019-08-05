package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGET is a function implements "/" path.
func IndexGET(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/browser")
}
