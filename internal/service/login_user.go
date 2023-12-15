package service

import (
	"BankApi/internal/domain"
	"BankApi/internal/service/utils"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type LoginUserUseCase struct {
	userRepository domain.UserRepository
	secretKey      string
}

func NewLoginUserUseCase(userRepository domain.UserRepository, secretKey string) *LoginUserUseCase {
	return &LoginUserUseCase{userRepository: userRepository, secretKey: secretKey}
}

type LoginUserCommand struct {
	Email    string
	Password []byte
}

func (useCase *LoginUserUseCase) Login(ctx context.Context, command LoginUserCommand) (string, error) {
	user, err := useCase.userRepository.FindByName(ctx, command.Email)
	if err != nil {
		return "", fmt.Errorf("find by username: %w", err)
	}

	log.Print("hash", user.PasswordHash())
	log.Println("password", command.Password)

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash(), command.Password); err != nil {
		return "", ErrNotCorrectCredentials
	}

	return utils.CreateToken(user, useCase.secretKey)

}

var ErrNotCorrectCredentials = errors.New("Not correct credentials")
