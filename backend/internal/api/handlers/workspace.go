package handlers

import (
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/models"
)

type WorkspaceHandler struct {
	DB (*sqlx.DB)
}

func NewWorkspaceHandler(db *sqlx.DB) *WorkspaceHandler {
	return &WorkspaceHandler{DB: db}
}

func (h *WorkspaceHandler) GetWorkspaces(c echo.Context) error {
	var workspaces []struct {
		Name        string `db:"name" json:"name"`
		Description string `db:"description" json:"description"`
	}

	sql, _, err := sq.Select("name", "description").From("workspaces").ToSql()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ToSQL for workspaces: %v", err)
	}

	err = h.DB.Select(&workspaces, sql)
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

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	workspace := models.Workspace{
		ID:          uuid.New(),
		Name:        input.Name,
		Description: input.Description,
	}

	sql, args, err := sq.
		Insert("workspaces").
		Columns("id", "name", "description").
		Values(workspace.ID, workspace.Name, workspace.Description).
		ToSql()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ToSQL for create Workspace: %v", err)
	}

	if _, err = h.DB.Exec(sql, args); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create workspace")
	}

	return c.JSON(http.StatusCreated, workspace)
}
