package auth

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	// テスト用の秘密鍵
	testSecret := "test-secret-key"
	os.Setenv("JWT_SECRET", testSecret)
	JwtSecret = []byte(testSecret)

	userID := "test-user-id"

	token, err := GenerateToken(userID)

	// エラーが無いことを確認
	assert.NoError(t, err)

	// トークンが空でないことを確認
	assert.NotEmpty(t, token)

	// トークンを解析して内容を確認
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
	
	// クレームの内容を確認
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, userID, claims["user_id"])	
}

func TestJwtCustomClaims(t *testing.T) {
	claims := JwtCustomClaims{
		UserID: "test-user-id",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	assert.Equal(t, "test-user-id", claims.UserID)
	assert.NotNil(t, claims.ExpiresAt)
}
