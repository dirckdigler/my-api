package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dirckdigler/my-api-golang/routers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", routers.GetTasks).Methods("GET")
	return router
}

func TestGetTask(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tasks", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
