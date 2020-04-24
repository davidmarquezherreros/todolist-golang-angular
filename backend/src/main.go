package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	parsedPage, err := strconv.Atoi(mux.Vars(r)["page"])
	if err == nil {
		json.NewEncoder(w).Encode(retrieveAllTasks(parsedPage))
	} else {
		RegisterLog(err.Error(), false)
	}
}

func updateTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func deleteTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Under development!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", serveIndex)

	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTaskById).Methods("GET")
	router.HandleFunc("/task/{id}", updateTaskById).Methods("PATCH")
	router.HandleFunc("/task/{id}", deleteTaskById).Methods("DELETE")
	router.HandleFunc("/tasks/{page}", getAllTasks).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))
}
