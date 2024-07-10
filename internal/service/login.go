package service

import (
	"borsodoy/radovid/pkg/utility"
	"net/http"
	"time"
)

type LoginProps struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	ExpireAt    time.Time `json:"expire_at"`
}

func Login(loginData LoginProps) (*LoginResponse, error) {
	if loginData.Email == "" || loginData.Password == "" {
		return nil, &utility.HttpError{Status: http.StatusBadRequest, Message: "Email and password are required."}
	}

	user, err := GetUserByEmail(loginData.Email)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			return nil, httpError
		}

		return nil, err
	}

	passwordAreEqual := utility.CompareHasAndPassword(loginData.Password, user.Password)

	if !passwordAreEqual {
		return nil, &utility.HttpError{Status: http.StatusUnauthorized, Message: "Wrong password"}
	}

	tokenData := utility.GenerateToken(user.ID, user.Email)

	loginResponse := &LoginResponse{
		AccessToken: tokenData.AccessToken,
		ExpireAt:    tokenData.ExpireAt,
	}

	return loginResponse, nil
}
