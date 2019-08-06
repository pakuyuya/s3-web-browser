package db


import (
    "database/sql"
    "fmt"

	_ "github.com/lib/pq"
	"s3-web-browser/server/go/setting"
)

// Connection is a function that get backend db
func Connection() (*sql.DB, error) {
	stg := setting.ServerSetting

	host := stg.DBHost
	port := stg.DBPort
	user := stg.DBUser
	pass := stg.DBPass
	dbname := stg.DBName

	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	return sql.Open("postgres", connstr)
}

// ConnectionForTest is a function that get test specified db
func ConnectionForTest() (*sql.DB, error) {
	host := "db"
	port := "5432"
	user := "postgres"
	pass := "passwrd"
	dbname := "s3webbrowser"

	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	return sql.Open("postgres", connstr)
}