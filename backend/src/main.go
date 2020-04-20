package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!") // TODO: Serve swagger API documentation
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(retrieveAllTasks())
}

func updateTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func deleteTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func main() {
	InitializeConfiguration()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", serveIndex)

	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTaskById).Methods("PATCH")
	router.HandleFunc("/tasks/{id}", deleteTaskById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9090", router))
}
