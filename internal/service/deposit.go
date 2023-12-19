package service

import "BankApi/internal/domain"

type DepositUseCase struct {
	billRepository domain.BillRepository
}

func NewDepositUseCase(billRepository domain.BillRepository) *DepositUseCase {
	return &DepositUseCase{billRepository: billRepository}
}
