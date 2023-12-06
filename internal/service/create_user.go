package service

import (
	"BankApi/internal/domain"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	user := domain.NewUser(command.Username, hash)

	err = useCase.userRepository.Save(ctx, user)
	if err != nil {
		return "", fmt.Errorf("save user: %w", err)
	}

	return useCase.createToken(user)
}

func (useCase *CreateUserUseCase) createToken(user *domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["user"] = user.Name()
	tokenString, err := token.SignedString([]byte(useCase.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

var ErrUnauthorized = errors.New("user is not authorized")

func (useCase *CreateUserUseCase) Login(ctx context.Context, command CreateUserCommand) (*domain.User, error) {
	user, err := useCase.userRepository.FindByName(ctx, command.Username)
	if err != nil {
		return nil, fmt.Errorf("find by username: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash(), command.Password); err != nil {
		return nil, ErrUnauthorized
	}

	return user, nil
}
