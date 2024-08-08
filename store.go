package main

import "database/sql"

type Store interface {
	CreateUser() error

	CreateTask(t *Task) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}

}

func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO  tasks (name , status , project_id , assigned_to) VALUES (?,?,?,?)", t.Name, t.Status, t.ProjectID, t.AssignedToId)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.ID = id
	return t, nil
}
