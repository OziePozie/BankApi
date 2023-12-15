package utils

import (
	"BankApi/internal/domain"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type Token string

type CreateTokenUseCase struct {
	secretKey string
}

func CreateToken(user *domain.User, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["user"] = user.Name()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	log.Print(tokenString)

	return tokenString, nil

}
