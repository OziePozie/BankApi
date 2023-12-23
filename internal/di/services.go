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
	getBill    *service.GetBillUseCase
	loginUser  *service.LoginUserUseCase
	deposit    *service.DepositUseCase
}

func (s *ServiceContainer) GetBill(ctx context.Context) *service.GetBillUseCase {

	if s.getBill == nil {
		s.getBill = service.NewGetBillUseCase(s.repo.BillRepository(ctx))
	}

	return s.getBill
}

func (s *ServiceContainer) LoginUser(ctx context.Context) *service.LoginUserUseCase {
	if s.loginUser == nil {
		s.loginUser = service.NewLoginUserUseCase(s.repo.UserRepository(ctx), s.secretKey)
	}

	return s.loginUser
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

func (s *ServiceContainer) Deposit(ctx context.Context) *service.DepositUseCase {

	if s.deposit == nil {
		s.deposit = service.NewDepositUseCase(s.repo.BillRepository(ctx), s.repo.TransactionManager(ctx))
	}
	return s.deposit
}
