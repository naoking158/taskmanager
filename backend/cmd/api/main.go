package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naoking158/taskmanager/internal/api"
	"github.com/naoking158/taskmanager/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("データベース接続の初期化に失敗しました: ", err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.SetupRoutes(e, db)

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
