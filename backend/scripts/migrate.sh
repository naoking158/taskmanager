#!/bin/bash

set -e

# データベース接続情報を環境変数から取得
DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# マイグレーションの実行
migrate -database "${DB_URL}" -path db/migrations "$@"
