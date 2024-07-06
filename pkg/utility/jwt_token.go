package utility

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("my_secret_key")

type GenerateTokenResponse struct {
  AccessToken string
  ExpireAt time.Time
}

func GenerateToken(userId uuid.UUID) (GenerateTokenResponse) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
    "user_id": userId,
		"expires_at": expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, _ := token.SignedString(jwtKey)

  return GenerateTokenResponse{
    AccessToken: tokenString,
    ExpireAt: expirationTime,
  }
}
