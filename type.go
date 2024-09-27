package main

import "time"

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	AssignedToID int64     `json:"assignedToID"`
	ProjectID    int64     `json:"projectId"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
