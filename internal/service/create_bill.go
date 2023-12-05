package service

import (
	"BankApi/internal/domain"
	"context"
	"github.com/gofrs/uuid"
)

type CreateBillUseCase struct {
	billRepository domain.BillRepository
}

func NewCreateBillUseCase(billRepository domain.BillRepository) *CreateBillUseCase {
	return &CreateBillUseCase{billRepository: billRepository}
}

type CreateCommand struct {
	Name   string
	UserID uuid.UUID
}

func (useCase *CreateBillUseCase) Handle(ctx context.Context, command CreateCommand) (*domain.Bill, error) {
	bill := domain.NewBill(command.Name, command.UserID)

	if err := bill.Validate(); err != nil {
		return nil, err
	}

	if err := useCase.billRepository.Save(ctx, bill); err != nil {
		return nil, err
	}

	return bill, nil
}
