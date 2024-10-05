package handlers

import (
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type UserInput struct {
	Username    string `json:"username" validate:"required,min=1,max=50"`
	DisplayName string `json:"display_name" validate:"max=50"`
}

type UserHandler struct {
	DB (*sqlx.DB)
	sq (*sq.StatementBuilderType)
}

func NewUserHandler(db *sqlx.DB, sq *sq.StatementBuilderType) *UserHandler {
	return &UserHandler{DB: db, sq: sq}
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	userID, err := uuid.Parse(c.Get("userID").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID")
	}

	var user struct {
		ID string `db:"id"`
		Username string `db:"username"`
		DisplayName string `db:"display_name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
	sql, args, err := h.sq.
		Select("id", "username", "display_name", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{"id": userID}).
		ToSql()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to ToSql for GetUserByID: %v", err)
	}

	err = h.DB.Get(&user, sql, args...)
	if err != nil {
		c.Logger().Errorf("Failed to get user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	userID, err := uuid.Parse(c.Get("userID").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID")
	}
	
	var input UserInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sql, args, err := h.sq.
		Update("users").
		SetMap(sq.Eq{
			"username": input.Username,
			"display_name": input.DisplayName,
		}).
		ToSql()
	if err != nil {
		 	return echo.NewHTTPError(http.StatusInternalServerError, "failed to ToSql for UpdateUser: %v", err)
	}

	if _, err = h.DB.Exec(sql, args...); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update user: %v", err)
	}

	return c.JSON(http.StatusOK, input)
}
