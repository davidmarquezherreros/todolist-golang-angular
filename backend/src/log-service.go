package main

import (
	"fmt"
	"os"
	"time"
)

func RegisterLog(error string, failedDb bool) {

	error = fmt.Sprintf("ERROR: %s				%s", error, time.Now().String())

	if configuration == nil {
		configuration = InitializeConfiguration()
	}

	if configuration.Logging.Database && !failedDb {
		parameters := []string{fmt.Sprintf("'%s'", error)}
		ExecuteSPWithParametersNoReturn("SP_INSERT_LOG_BACKEND", parameters)
	} else {
		f, _ := os.OpenFile(configuration.Logging.Filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
		f.WriteString(error)
		f.Close()
	}
}
