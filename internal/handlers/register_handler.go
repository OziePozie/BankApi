package handlers

import "BankApi/internal/service"

type POSTRegisterHandler struct {
	useCase *service.CreateUserUseCase
}

func NewPOSTRegisterHandler(useCase *service.CreateUserUseCase) *POSTRegisterHandler {
	return &POSTRegisterHandler{useCase: useCase}
}

type POSTRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
