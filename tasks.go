package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// all the service use store/repository

type TaskService struct {
	store Store
}

func NewTaskSerive(s Store) *TaskService {
	return &TaskService{
		store: s,
	}
}

func (s *TaskService) RegisterRouter(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	defer r.Body.Close()

	var task *Task
	if err := json.Unmarshal(body, &task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{
			"Invalid payload",
		})
	}

	if err := validateTask(task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{
			err.Error(),
		})
		return
	}

	task, err = s.store.CreateTask(task)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	WriteJson(w, http.StatusCreated, task)

}

func (s *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTask(task *Task) error {
	if len(strings.TrimSpace(task.Name)) == 0 {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIdRequired
	}

	if task.AssignedToID == 0 {
		return errAssignedRequired
	}

	return nil
}

var errNameRequired = errors.New("name is required")

var errProjectIdRequired = errors.New("project id is required")

var errAssignedRequired = errors.New("user id is required")
