package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/models"
)

type TaskHandler struct {
	DB (*sqlx.DB)
}

func NewTaskHandler(db *sqlx.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	taskID := c.Param("taskID")
	if taskID == "" {
		return echo.NewHTTPError(http.StatusBadGateway, "Invalid endpoint")
	}

	var task models.Task
	query := "SELECT * FROM tasks WHERE id = $1"
	err := h.DB.Get(&task, query, taskID)
	if err != nil {
		c.Logger().Errorf("Failed to get tasks: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get task")
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetTaskAll(c echo.Context) error {
	workspaceID := c.Param("workspaceID")
	if workspaceID == "" {
		return echo.NewHTTPError(http.StatusBadGateway, "Invalid endpoint")
	}

	var tasks []struct {
		Id      uuid.UUID `json:"id" db:"id"`
		Title   string    `db:"title" json:"title"`
		Status  string    `db:"status" json:"status"`
		DueDate time.Time `json:"due_date" db:"due_date"`
	}

	query := "SELECT id, title, status, due_date FROM tasks WHERE workspace_id = $1"
	err := h.DB.Select(&tasks, query, workspaceID)
	if err != nil {
		c.Logger().Errorf("Failed to get tasks: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tasks")
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusOK, []models.Task{})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	workspaceID, err := uuid.Parse(c.Param("workspaceID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid workspace ID")
	}

	var input struct {
		Title       string            `json:"title" validate:"required"`
		Description string            `json:"description"`
		Status      models.TaskStatus `json:"status" validate:"required"`
		AssignedTo  string            `json:"assigned_to"`
		DueDate     time.Time         `json:"due_date"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !input.Status.IsValid() {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status")
	}

	userID, err := uuid.Parse(c.Get("userID").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID")
	}

	var assignedTo *uuid.UUID
	if input.AssignedTo != "" {
		pUUID, err := uuid.Parse(input.AssignedTo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid assigned to ID")
		}
		assignedTo = (&pUUID)
	}

	task := &models.Task{
		ID:          uuid.New(),
		WorkspaceID: workspaceID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedBy:   userID,
		AssignedTo:  assignedTo,
		DueDate:     input.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `
        INSERT INTO tasks (id, workspace_id, title, description, status, created_by, assigned_to, due_date, created_at, updated_at)
        VALUES (:id, :workspace_id, :title, :description, :status, :created_by, :assigned_to, :due_date, :created_at, :updated_at)
    `

	if _, err := h.DB.NamedExec(query, task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create task")
	}

	return c.JSON(http.StatusCreated, task)
}


func (h *TaskHandler) EditTask(c echo.Context) error {
	workspaceID, err := uuid.Parse(c.Param("workspaceID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid workspace ID")
	}

	var input struct {
		ID uuid.UUID `json:"id" validate:"required"`
		Title       string            `json:"title" validate:"required"`
		Description string            `json:"description"`
		Status      models.TaskStatus `json:"status" validate:"required"`
		AssignedTo  string            `json:"assigned_to"`
		DueDate     time.Time         `json:"due_date"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if !input.Status.IsValid() {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status")
	}

	userID, err := uuid.Parse(c.Get("userID").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID")
	}


	var assignedTo *uuid.UUID
	if input.AssignedTo != "" {
		pUUID, err := uuid.Parse(input.AssignedTo)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid assigned to ID")
		}
		assignedTo = (&pUUID)
	}

	task := &models.Task{
		ID:          input.ID,
		WorkspaceID: workspaceID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		AssignedTo:  assignedTo,
		DueDate:     input.DueDate,
		UpdatedAt:   time.Now(),
	}

	query := `
        UPDATE tasks
   SET workspace_id = :workspace_id,
       title = :title,
       description = :description,
       status = :status,
       assigned_to = :assigned_to,
       due_date = :due_date,
       updated_at = :updated_at
 WHERE id = :id
    `

	if _, err := h.DB.NamedExec(query, task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update task: %w", err)
	}

	return c.JSON(http.StatusCreated, task)
}
