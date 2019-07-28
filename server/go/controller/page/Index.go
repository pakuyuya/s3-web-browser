package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
