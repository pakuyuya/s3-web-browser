package db


import (
    "database/sql"
    "fmt"

	_ "github.com/lib/pq"
	"s3-web-brawser/server/go/setting"
)

// Connection is a function that get backend db
func Connection() (*sql.Conn, error) {
	host := setting.DBHost
	port := setting.DBPort
	user := setting.DBUser
	pass := setting.DBPass
	dbname := setting.DBPName

	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	return sql.Open("postgres", connstr)
}