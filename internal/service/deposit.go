package service

import (
	"BankApi/internal/domain"
	"BankApi/internal/pkg"
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"log"
)

type DepositUseCase struct {
	billRepository     domain.BillRepository
	transactionManager pkg.TransactionManager
}

type DepositCommand struct {
	UserID uuid.UUID
	BillID uuid.UUID
	Amount int
}

func NewDepositUseCase(billRepository domain.BillRepository, manager pkg.TransactionManager) *DepositUseCase {
	return &DepositUseCase{billRepository: billRepository,
		transactionManager: manager}
}

func (c *DepositUseCase) Handle(ctx context.Context, command DepositCommand) (int, error) {

	var newBalance int
	log.Print(c.billRepository)
	log.Print(c.transactionManager)
	err := c.transactionManager.Do(ctx, func(ctx context.Context) error {
		log.Print("billRepos:=====", c.billRepository)
		bill, err := c.billRepository.GetBillByBillIDAndUserIDEquals(ctx, command.UserID, command.BillID)
		if err != nil {
			return err
		}
		//if !bill.IsOpen() {
		//	return errors.New("bill is closed")
		//}
		if bill.Balance() < command.Amount {
			return errors.New("balance is less than amount")
		}
		newBalance, err = c.billRepository.DepositAmount(ctx, command.BillID, command.Amount)
		if err != nil {
			return err
		}
		log.Print(err)

		return err
	})
	if err != nil {
		return newBalance, err
	}

	return newBalance, err
}
