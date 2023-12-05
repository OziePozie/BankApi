package handlers

import (
	"BankApi/internal/service"
	"net/http"
)

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

type POSTRegisterResponse struct {
	Token string `json:"token"`
}

func (P POSTRegisterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
