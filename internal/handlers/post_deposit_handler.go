package handlers

import "BankApi/internal/service"

type POSTDepositHandler struct {
	useCase *service.DepositUseCase
}

func NewPOSTDepositHandler(useCase *service.DepositUseCase) *POSTDepositHandler {
	return &POSTDepositHandler{useCase: useCase}
}
