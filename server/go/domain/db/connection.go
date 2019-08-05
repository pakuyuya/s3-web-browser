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

// ConnectionForTest is a function that get test specified db
func ConnectionForTest() (*sql.Conn, error) {
	host := "db"
	port := "5432"
	user := "postgres"
	pass := "passwrd"
	dbname := "s3webbrowser"

	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	return sql.Open("postgres", connstr)
}