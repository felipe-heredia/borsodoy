package utility

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("my_secret_key")

type GenerateTokenResponse struct {
	AccessToken string
	ExpireAt    time.Time
}

type JWTCustomClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	Email     string    `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uuid.UUID, email string) GenerateTokenResponse {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := JWTCustomClaims{
		UserID:    userId,
		ExpiresAt: expirationTime,
		Email:     email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	return GenerateTokenResponse{
		AccessToken: tokenString,
		ExpireAt:    expirationTime,
	}
}

type validateTokenResponse struct {
	Token  *jwt.Token
	Cliams JWTCustomClaims
}

func ValidateToken(tokenString string) (*validateTokenResponse, error) {
	claims := &JWTCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, &HttpError{Status: http.StatusUnauthorized, Message: "Invalid token"}
	}

	response := &validateTokenResponse{
		Token:  token,
		Cliams: *claims,
	}

	return response, nil
}
