package di

import (
	"BankApi/internal/service"
	"os"
)

type ServiceContainer struct {
	repo *RepoContainer

	secretKey string

	createBill *service.CreateBillUseCase
	createUser *service.CreateUserUseCase
}

func (s *ServiceContainer) SecretKey() string {
	if s.secretKey == "" {
		s.secretKey = os.Getenv("JWT_SECRET_KEY")
	}

	return s.secretKey
}

func (s *ServiceContainer) CreateBill() *service.CreateBillUseCase {

	if s.createBill == nil {
		s.createBill = service.NewCreateBillUseCase(s.repo.billRepository)
	}
	return s.createBill
}

func (s *ServiceContainer) SetCreateBill(createBill *service.CreateBillUseCase) {
	s.createBill = createBill
}

func (s *ServiceContainer) CreateUser() *service.CreateUserUseCase {

	if s.createUser == nil {
		s.createUser = service.NewCreateUserUseCase(s.repo.userRepository, s.secretKey)
	}

	return s.createUser
}

func (s *ServiceContainer) SetCreateUser(createUser *service.CreateUserUseCase) {
	s.createUser = createUser
}

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{}
}
