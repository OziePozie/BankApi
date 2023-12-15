package service

import (
	"BankApi/internal/domain"
	"BankApi/internal/service/utils"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepository domain.UserRepository
	secretKey      string
}

func NewCreateUserUseCase(userRepository domain.UserRepository, secretKey string) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository, secretKey: secretKey}
}

type CreateUserCommand struct {
	Username string
	Email    string
	Password []byte
}

func (useCase *CreateUserUseCase) Register(ctx context.Context, command CreateUserCommand) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(command.Password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := domain.NewUser(command.Username, command.Email, hash)

	err = useCase.userRepository.Save(ctx, user)
	if err != nil {
		return "", fmt.Errorf("save user: %w", err)
	}

	return utils.CreateToken(user, useCase.secretKey)
}

//
//func (useCase *CreateUserUseCase) createToken(user *domain.User) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//	claims["exp"] = time.Now().Add(10 * time.Minute)
//	claims["user"] = user.Name()
//	tokenString, err := token.SignedString([]byte(useCase.secretKey))
//	if err != nil {
//		return "", err
//	}
//
//	log.Print(tokenString)
//
//	return tokenString, nil
//}

var ErrUnauthorized = errors.New("user is not authorized")
