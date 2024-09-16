package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

type Task struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	WorkspaceID  uuid.UUID  `json:"workspace_id" db:"workspace_id"`
	Title        string     `json:"title" db:"title"`
	Description  string     `json:"description" db:"description"`
	Status       TaskStatus `json:"status" db:"status"`
	CreatedBy    uuid.UUID  `json:"created_by" db:"created_by"`
	AssignedTo   *uuid.UUID `json:"assigned_to" db:"assigned_to"`
	ParentTaskID *uuid.UUID `db:"parent_task_id" json:"parent_task_id"`
	DueDate      time.Time  `json:"due_date" db:"due_date"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "In Progress"
	TaskStatusDone       TaskStatus = "DONE"
	TaskStatusOnHold     TaskStatus = "On Hold"
)

func (s TaskStatus) IsValid() bool {
	switch s {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone, TaskStatusOnHold:
		return true
	}
	return false
}
