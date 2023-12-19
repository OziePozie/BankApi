package service

import (
	"BankApi/internal/domain"
	"BankApi/internal/handlers/middleware"
	"context"
)

type GetBillUseCase struct {
	billRepository domain.BillRepository
}

func NewGetBillUseCase(billRepository domain.BillRepository) *GetBillUseCase {
	return &GetBillUseCase{billRepository: billRepository}
}

type GetBillCommand struct {
	Page  int
	Items int
	Name  string
}

func (useCase *GetBillUseCase) Handle(ctx context.Context, command GetBillCommand) (*[]domain.Bill, error) {
	bills, err := useCase.billRepository.FindAllByUserWithILIKE(ctx, command.Name, command.Page, command.Items,
		middleware.UserIDFromContext(ctx))
	if err != nil {
		return nil, err
	}
	return bills, nil
}
