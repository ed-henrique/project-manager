package models

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          int
	Name        string
	Description string
	Project     int
	Priority    int
	Deadline    time.Time
}

type TaskModel struct {
	db *sql.DB
}

func NewTaskModel(db *sql.DB) *ProjectModel {
	return &ProjectModel{db: db}
}

func (tm *TaskModel) Create(name, description string, project, priority int, deadline time.Time) error {
	_, err := tm.db.Exec("INSERT INTO TASK (NAME, DESCRIPTION, PROJECT, PRIORITY) VALUES (?, ?, ?, ?)",
		name,
		description,
		project,
		priority,
		deadline,
	)
	return err
}
