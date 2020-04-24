package main

import (
	"reflect"
	"strconv"
)

func retrieveAllTasks(pageNumber int) []Task {
	start := (pageNumber - 1) * 20
	end := pageNumber * 20

	rows := ExecuteSPWithParametersWithReturn("SP_RETRIEVE_ALL_TASKS", []string{strconv.Itoa(start), strconv.Itoa(end)})

	var tasks []Task

	for rows.Next() {
		task := Task{}

		s := reflect.ValueOf(&task).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			RegisterLog(err.Error(), false)
		}
		tasks = append(tasks, task)
	}
	return tasks
}
