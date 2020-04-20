package main

import (
	"reflect"
)

func retrieveAllTasks() []Task {
	rows := ExecuteSP("SP_RETRIEVE_ALL_TASKS")

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
			//TODO: LOG
		}
		tasks = append(tasks, task)
	}
	return tasks
}
