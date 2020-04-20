package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID         int
	NAME       string
	STATUS     string
	DATE_START string
	DATE_END   string
}

var db *sql.DB
var err error

///
/// Pings the database
///
func Ping() bool {
	err = db.Ping()
	if err != nil {
		return false
	} else {
		return true
	}
}

///
/// Executes a defined stored procedure
///
func ExecuteSP(spname string) *sql.Rows {

	db, err = sql.Open("mysql", GetConnectionString())

	var rows *sql.Rows

	rows, err := db.Query(fmt.Sprintf("CALL `%s`()", spname))

	if err != nil {
		//TODO: Log error
	}

	defer db.Close()
	return rows
}
