package service

import (
	"BankApi/internal/domain"
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
	Password []byte
}

func (useCase *CreateUserUseCase) Register(command CreateUserCommand) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(command.Password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := domain.NewUser(command.Username, hash)

	err = useCase.userRepository.Save(user)
	if err != nil {
		return "", fmt.Errorf("save user: %w", err)
	}

	return useCase.createToken(user)
}

func (useCase *CreateUserUseCase) createToken(user *domain.User) (string, error) {
	return "", errors.New("not implemented")
}

var ErrUnauthorized = errors.New("user is not authorized")

func (useCase *CreateUserUseCase) Login(command CreateUserCommand) (*domain.User, error) {
	user, err := useCase.userRepository.FindByName(command.Username)
	if err != nil {
		return nil, fmt.Errorf("find by username: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash(), command.Password); err != nil {
		return nil, ErrUnauthorized
	}

	return user, nil
}
