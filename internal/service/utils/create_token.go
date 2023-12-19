package utils

import (
	"BankApi/internal/domain"
	"github.com/golang-jwt/jwt/v5"
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
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(10 * time.Hour))
	claims["user"] = user.ID()
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	log.Print(tokenString)

	return tokenString, nil

}
