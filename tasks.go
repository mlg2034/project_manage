package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var errNameRequired = errors.New("name is required")
var errProjectIdRequired = errors.New("projectId is required")
var errUserIDRequired = errors.New("userId is required")

type TaskService struct {
	store Store
}

func NewTaskService(store Store) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/tasks/", s.handleCreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")

}

func (s *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{
			Error: "Invalid request payload",
		})
		return
	}

	defer r.Body.Close()
	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{
			Error: "Invalid request payload",
		})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	t, err := s.store.CreateTask(task)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to create task",
		})
		return
	}

	WriteJson(w, http.StatusCreated, t)
}

func (s *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}
	if task.ProjectID == 0 {
		return errProjectIdRequired
	}
	if task.AssignedToId == 0 {
		return errUserIDRequired
	}
	return nil
}
