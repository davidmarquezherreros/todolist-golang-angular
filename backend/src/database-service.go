package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int
	NAME        string
	DESCRIPTION string
	STATUS      string
	DATE_START  string
	DATE_END    string
}

var db *sql.DB
var err error

///
/// Pings the database
///
func Ping() bool {
	err = db.Ping()
	if err != nil {
		RegisterLog(err.Error(), true)
		return false
	} else {
		return true
	}
}

///
/// Executes a defined stored procedure without parameters
///
func ExecuteSP(spname string) *sql.Rows {

	db, err = sql.Open("mysql", GetConnectionString())

	var rows *sql.Rows

	rows, err := db.Query(fmt.Sprintf("CALL `%s`()", spname))

	if err != nil {
		RegisterLog(err.Error(), true)
	}

	defer db.Close()
	return rows
}

///
/// Executes a defined stored procedure with parameters
/// Does not return records
///
func ExecuteSPWithParametersNoReturn(spname string, parameters []string) {

	db, err = sql.Open("mysql", GetConnectionString())

	_, err := db.Exec(BuildSqlStatement(spname, parameters))

	if err != nil {
		RegisterLog(err.Error(), true)
	}

	defer db.Close()
}

///
/// Executes a defined stored procedure with parameters
/// Does return records
///
func ExecuteSPWithParametersWithReturn(spname string, parameters []string) *sql.Rows {

	db, err = sql.Open("mysql", GetConnectionString())

	rows, err := db.Query(BuildSqlStatement(spname, parameters))

	if err != nil {
		RegisterLog(err.Error(), true)
	}

	defer db.Close()
	return rows
}

func BuildSqlStatement(spname string, parameters []string) string {

	return fmt.Sprintf("CALL `%s`( %s )", spname, strings.Join(parameters, ","))
}
