package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/models"
)

type CommentHandler struct {
	DB       (*sqlx.DB)
}

func NewCommentHandler(db *sqlx.DB) *CommentHandler {
	return &CommentHandler{DB: db}
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var input struct {
		Content string `json:"content" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userID, err := uuid.Parse(c.Get("userID").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID")
	}

	comment := &models.Comment{
		ID:        uuid.New(),
		TaskID:    taskID,
		UserID:    userID,
		Content:   input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
        INSERT INTO comments (id, task_id, user_id, content, created_at, updated_at)
        VALUES (:id, :task_id, :user_id, :content, :created_at, :updated_at)
    `
	if _, err = h.DB.NamedExec(query, comment); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create comment: ", err)
	}

	return c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) GetCommentsByTaskID(c echo.Context) error {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	query := `
        SELECT id, task_id, user_id, content, created_at, updated_at
        FROM comments
        WHERE task_id = $1
        ORDER BY created_at DESC
    `

	var comments []models.Comment
	err = h.DB.Select(&comments, query, taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get comments: ", err)
	}

	return c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) UpdateComment(c echo.Context) error {
	commentID, err := uuid.Parse(c.Param("commentID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid comment ID")
	}

	var input struct {
		Content string `json:"content" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment := &models.Comment{
		ID:        commentID,
		Content:   input.Content,
		UpdatedAt: time.Now(),
	}

	query := `
        UPDATE comments
        SET content = $1, updated_at = $2
        WHERE id = $3
    `
	if _, err = h.DB.Exec(query, comment.Content, comment.UpdatedAt, comment.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update comment: ", err)
	}

	return c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	commentID, err := uuid.Parse(c.Param("commentID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid comment ID")
	}

	query := "DELETE FROM comments WHERE id = $1"
	if _, err := h.DB.Exec(query, commentID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete comment")
	}

	return c.NoContent(http.StatusNoContent)
}
