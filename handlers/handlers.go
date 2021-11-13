package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dirckdigler/my-api-golang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// HandlersRoute sets port and handle routes.
// @host localhost:8081
// @BasePath /tasks
// @BasePath /task/{id}
// @BasePath /create-task
// @BasePath /update-task/{id}
// @BasePath /delete-task/{id}
func HandlersRoute() {
	router := mux.NewRouter()
	router.HandleFunc("/", routers.IndexRoute).Methods("GET")
	router.HandleFunc("/tasks", routers.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", routers.GetTaskByID).Methods("GET")
	router.HandleFunc("/create-task", routers.CreateTask).Methods("POST")
	router.HandleFunc("/update-task/{id}", routers.UpdateTask).Methods("PUT")
	router.HandleFunc("/delete-task/{id}", routers.DeleteTask).Methods("DELETE")

	PORT := os.Getenv("PORT")
	// Validate if the environment variable does not exist yet.
	if PORT == "" {
		PORT = "8081"
	}
	// Object cors are the permissions given to the API to make it accessible from anywhere.
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
