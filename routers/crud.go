package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dirckdigler/my-api-golang/models"
	"github.com/gorilla/mux"
)

type allTasks []models.Person

var tasks = allTasks{
	{
		ID:        1,
		FirstName: "Task One",
		Lastname:  "Some Content",
	},
}

// GetTasks get all tasks.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID get specific task depending ID parameter.
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Get parameter.
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, value := range tasks {
		if value.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(value)
		}
	}

}

// CreateTask create task.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Person
	// ioutil manage enter and output server.
	// Get request information.
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Soemthing went wrong...Insert valid person data")
	}
	// Assing this information to newTask variable.
	json.Unmarshal(reqBody, &newTask)
	// Create an ID dinammically.
	newTask.ID = len(tasks) + 1
	// Save this new information in list of tasks.
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Respond client with the information.
	json.NewEncoder(w).Encode(newTask)
}

// UpdateTask update specific task depending ID parameter.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Person

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = t.ID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
		}
	}

}

// DeleteTask delete specific task depending ID parameter.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Get parameter.
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, value := range tasks {
		if value.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task ID %v was removed succesfully", taskID)
		}
	}
}
