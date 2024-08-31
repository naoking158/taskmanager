package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/naoking158/taskmanager/internal/api/handlers"
	"github.com/naoking158/taskmanager/internal/middleware"
)

func SetupRoutes(e *echo.Echo, db *sqlx.DB) {
	authHandler := handlers.NewAuthHandler(db)

	v1 := e.Group("/api/v1")

	// 認証不要
	auth:= v1.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	// 認証必要
	authenticated := v1.Group("")
	authenticated.Use(middleware.JWTMiddleware())
	authenticated.Use(middleware.Auth)

	workspaceHandler := handlers.NewWorkspaceHandler(db)
	authenticated.GET("/workspaces", workspaceHandler.GetWorkspaces)
	authenticated.POST("/workspaces", workspaceHandler.CreateWorkspace)
}
