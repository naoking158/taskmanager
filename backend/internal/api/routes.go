package api

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/api/handlers"
	"github.com/naoking158/taskmanager/internal/middleware"
)

func SetupRoutes(e *echo.Echo, db *sqlx.DB, sq *squirrel.StatementBuilderType) {
	authHandler := handlers.NewAuthHandler(db, sq)

	v1 := e.Group("/api/v1")

	// 認証不要
	auth := v1.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	// 認証必要
	authenticated := v1.Group("")
	authenticated.Use(middleware.JWTMiddleware())
	authenticated.Use(middleware.Auth)

	workspaceHandler := handlers.NewWorkspaceHandler(db, sq)
	authenticated.GET("/workspaces", workspaceHandler.GetWorkspaces)
	authenticated.POST("/workspaces", workspaceHandler.CreateWorkspace)

	taskHandler := handlers.NewTaskHandler(db, sq)
	authenticated.GET("/workspaces/:workspaceID/tasks", taskHandler.GetTaskAll)
	authenticated.POST("/workspaces/:workspaceID/tasks", taskHandler.CreateTask)
	authenticated.GET("/tasks/:taskID", taskHandler.GetTask)

	commentHandler := handlers.NewCommentHandler(db)
	authenticated.POST("/tasks/:taskID/comments", commentHandler.CreateComment)
	authenticated.GET("/tasks/:taskID/comments", commentHandler.GetCommentsByTaskID)
	authenticated.PUT("/comments/:commentID", commentHandler.UpdateComment)
	authenticated.DELETE("/comments/:commentID", commentHandler.DeleteComment)
}
