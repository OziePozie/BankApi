package handlers

import (
	"BankApi/internal/service"
	"encoding/json"
	"log"
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
	Email    string `json:"email"`
	Password string `json:"password"`
}

type POSTRegisterResponse struct {
	Token []byte `json:"token"`
}

func (p *POSTRegisterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var r POSTRegisterRequest
	if err := json.NewDecoder(request.Body).Decode(&r); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	log.Print(r.Email)

	token, err := p.useCase.Register(
		request.Context(),
		service.CreateUserCommand{
			Username: r.Username,
			Email:    r.Email,
			Password: []byte(r.Password),
		},
	)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}
	writer.Header().Set("Authorization", token)
	writer.WriteHeader(http.StatusOK)

}
