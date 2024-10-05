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
	sq (*sq.StatementBuilderType)
}

func NewWorkspaceHandler(db *sqlx.DB, sq *sq.StatementBuilderType) *WorkspaceHandler {
	return &WorkspaceHandler{DB: db, sq: sq}
}

func (h *WorkspaceHandler) GetWorkspaces(c echo.Context) error {
	var workspaces []struct {
		ID          string `db:"id" json:"id"`
		Name        string `db:"name" json:"name"`
		Description string `db:"description" json:"description"`
	}

	sql, _, err := h.sq.Select("id", "name", "description").From("workspaces").ToSql()
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

	sql, args, err := h.sq.
		Insert("workspaces").
		Columns("id", "name", "description").
		Values(workspace.ID, workspace.Name, workspace.Description).
		ToSql()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ToSQL for create Workspace: %v", err)
	}

	if _, err = h.DB.Exec(sql, args...); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create workspace")
	}

	return c.JSON(http.StatusCreated, workspace)
}

func (h *WorkspaceHandler) DeleteWorkspace(c echo.Context) error {
	workspaceID := c.Param("workspaceID")
	if workspaceID == "" {
		return echo.NewHTTPError(http.StatusBadGateway, "Invalid endpoint")
	}

	sql, args, err := h.sq.Delete("workspaces").Where(sq.Eq{"id": workspaceID}).ToSql()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ToSql for DeleteWorkspace: %v", err)
	}

	if _, err = h.DB.Exec(sql, args...); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete workspace: %v", err)
	}

	return c.JSON(http.StatusNoContent, nil)
}
