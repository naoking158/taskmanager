package handlers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/models"
)

type WorkspaceHandler struct {
	DB       *sqlx.DB
	Validate (*validator.Validate)
}

func NewWorkspaceHandler(db *sqlx.DB) *WorkspaceHandler {
	return &WorkspaceHandler{
		DB: db,
		Validate: validator.New(),
	}
}

func (h *WorkspaceHandler) GetWorkspaces(c echo.Context) error {
	var workspaces []struct {
		Name        string `db:"name" json:"name"`
		Description string `db:"description" json:"description"`
	}

	err := h.DB.Select(&workspaces, "SELECT name, description FROM workspaces")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch workspaces")
	}

	if len(workspaces) == 0 {
		log.Printf("No workspaces found")
	}
	
	return c.JSON(http.StatusOK, workspaces)
}

func (h *WorkspaceHandler) CreateWorkspace(c echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Validate.Struct(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
		
	workspace := models.Workspace{
		ID:          uuid.New(),
		Name:        input.Name,
		Description: input.Description,
	}

	_, err := h.DB.NamedExec(`
        INSERT INTO workspaces (id, name, description)
        VALUES (:id, :name, :description)
    `, workspace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create workspace")
	}

	return c.JSON(http.StatusCreated, workspace)
}
