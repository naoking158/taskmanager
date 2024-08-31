package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB       (*sqlx.DB)
	Validate (*validator.Validate)
}

type RegisterInput struct {
	Username    string `json:"username" validate:"required,min=3,max=50"`
	Password    string `json:"password" validate:"required,min=8"`
	DisplayName string `json:"display_name" validate:"max=100"`
}

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewAuthHandler(db *sqlx.DB) *AuthHandler {
	return &AuthHandler{
		DB:       db,
		Validate: validator.New(),
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var input RegisterInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "無効な入力です", err)
	}

	if err := h.Validate.Struct(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// トランザクションの開始
	tx, err := h.DB.Beginx()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "トランザクションの開始に失敗しました: ", err)
	}
	defer tx.Rollback()

	// ユーザー名の重複チェック
	var count int
	err = tx.Get(&count, "SELECT COUNT(username) FROM users WHERE username = $1", input.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザー名の確認中にエラーが発生しました: ", err)
	}
	if count > 0 {
		return echo.NewHTTPError(http.StatusConflict, "username already exists")
	}

	// パスワードのハッシュ化
	hPW, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "パスワードのハッシュ化に失敗しました: ", err)
	}

	// 新しいユーザーの作成
	newUser := struct {
		ID           uuid.UUID `db:"id"`
		Username     string    `db:"username"`
		PasswordHash string    `db:"password_hash"`
		DisplayName  string    `db:"display_name"`
	}{
		ID:           uuid.New(),
		Username:     input.Username,
		PasswordHash: string(hPW),
		DisplayName:  input.DisplayName,
	}

	// ユーザーをデータベースに挿入
	_, err = tx.NamedExec(`
													INSERT INTO users (id, username, password_hash, display_name)
													VALUES (:id, :username, :password_hash, :display_name)
													`, newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ユーザーの作成に失敗しました: ", err)
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "トランザクションのコミットに失敗しました: ", err)
	}

	return c.JSON(http.StatusCreated, newUser)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var input LoginInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "無効な入力です: ", err)
	}

	if err := h.Validate.Struct(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// データベースからユーザーを取得
	var user struct {
		ID           string `db:"id"`
		PasswordHash string `db:"password_hash"`
	}
	err := h.DB.Get(&user, "SELECT id, password_hash FROM users WHERE username = $1", input.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found.")
	}

	// パスワードチェック
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password.")
	}

	// JWT トークンを生成
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "トークンの生成に失敗しました: ", err)
	}

	return c.JSON(http.StatusOK, token)
}
