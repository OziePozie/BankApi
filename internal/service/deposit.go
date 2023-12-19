package service

import (
	"BankApi/internal/domain"
	"context"
	"github.com/gofrs/uuid"
	"log"
)

type DepositUseCase struct {
	billRepository domain.BillRepository
}

type DepositCommand struct {
	UserID uuid.UUID
	Amount int
}

func NewDepositUseCase(billRepository domain.BillRepository) *DepositUseCase {
	return &DepositUseCase{billRepository: billRepository}
}

func (c DepositUseCase) Handle(ctx context.Context, command DepositCommand) (int, error) {
	bill, err := c.billRepository.DepositAmount(ctx, command.UserID, command.Amount)

	log.Print("Ошибка в депозите", err)

	log.Print("New balance", bill.Balance())

	return bill.Balance(), err
}
