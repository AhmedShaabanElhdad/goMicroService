package main

import "database/sql"

// the value of make it interface
// - it is easily to inject it
// - it is easily to test it
// - easily to mock it
type Store interface {

	// User
	CreateUser() error

	// Task
	CreateTask(tast *Task) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s Storage) CreateUser() error {
	return nil
}

func (s Storage) CreateTask(task *Task) (*Task, error) {
	rows, err := s.db.Exec(`
	Insert into tasks(
		name ,status,assignedToID,projectId 
		) value(?,?,?,?)
	`, task.Name, task.Status, task.AssignedToID, task.ProjectID)

	if err != nil {
		return nil, err
	}

	Id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.ID = Id

	return task, nil
}
