package di

import (
	"BankApi/internal/service"
	"context"
	"os"
)

type ServiceContainer struct {
	repo *RepoContainer

	secretKey string

	createBill *service.CreateBillUseCase
	createUser *service.CreateUserUseCase
}

func NewServiceContainer(repo *RepoContainer) *ServiceContainer {
	return &ServiceContainer{repo: repo}
}

func (s *ServiceContainer) SecretKey() string {
	if s.secretKey == "" {
		s.secretKey = os.Getenv("JWT_SECRET_KEY")
	}

	return s.secretKey
}

func (s *ServiceContainer) CreateBill(ctx context.Context) *service.CreateBillUseCase {

	if s.createBill == nil {
		s.createBill = service.NewCreateBillUseCase(s.repo.BillRepository(ctx))
	}
	return s.createBill
}

func (s *ServiceContainer) SetCreateBill(createBill *service.CreateBillUseCase) {
	s.createBill = createBill
}

func (s *ServiceContainer) CreateUser(ctx context.Context) *service.CreateUserUseCase {

	if s.createUser == nil {
		s.createUser = service.NewCreateUserUseCase(s.repo.UserRepository(ctx), s.secretKey)
	}

	return s.createUser
}

func (s *ServiceContainer) SetCreateUser(createUser *service.CreateUserUseCase) {
	s.createUser = createUser
}
