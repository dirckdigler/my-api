package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	ID        int    `json:"ID,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
}

// Persistence
var tasks = allTasks{
	{
		ID:        1,
		FirstName: "Task One",
		Lastname:  "Some Content",
	},
}

type allTasks []person

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask person
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Esta mal")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API 2")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/task", getTask).Methods("GET")
	router.HandleFunc("/create-task", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
