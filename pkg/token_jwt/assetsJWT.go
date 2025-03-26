package tokenjwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var SecretKey string

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT создаёт и подписывает JWT токен для указанного пользователя.
func GenerateJWT(userID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey)) // ✅ используем []byte
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	fmt.Println("Generated Token:", signedToken)
	return signedToken, nil
}

// DecodeJWT парсит и валидирует токен, возвращает Claims если всё ок.
func DecodeJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // ✅ правильно: []byte, а не string
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// InitJWTKey инициализирует секретный ключ для JWT.
func InitJWTKey(key string) {
	SecretKey = key
	fmt.Println("✅ JWT key initialized successfully")
}
