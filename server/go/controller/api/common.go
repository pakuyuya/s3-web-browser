package api
import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func responseError(c *gin.Context, errorcode int, msg string) {
	c.JSON(errorcode, gin.H{
		"result": "error",
		"message": msg,
	})
}

func getConnTx() (*sql.Database, *sql.Tx, error) {
	conn, err := db.Connection()
	if err != nil {
		return nil, nil, err
	}
	tx, err := conn.Begin()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}
	return conn, tx, nil
}