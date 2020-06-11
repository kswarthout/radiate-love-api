package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kswarthout/radiate-love-api/domain"
)

// AuthService provides authenticates users and handles JWT tokens
type AuthService struct {
	Config *domain.Config
}

// GenerateJWT generates a JWT token from user credentials
func (a *AuthService) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Rick Sanchez"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// tokenString, err := token.SignedString(config.Constants.JwtSigningKey)
	tokenString, err := token.SignedString("")
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
