package service

import (
	"BankApi/internal/domain"
	"BankApi/internal/service/utils"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash(), command.Password); err != nil {
		return "", ErrUnauthorized
	}

	return utils.CreateToken(user, useCase.secretKey)

}
