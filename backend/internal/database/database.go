package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase() (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
	  return nil, fmt.Errorf("データベース接続に失敗しました: %v", err)
	}

	// コネクションプールの設定
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	// 接続テスト
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("データベースの接続テストに失敗しました: %v", err)
	}

	return db, nil
}
