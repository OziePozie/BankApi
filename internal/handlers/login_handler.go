package handlers

import (
	"BankApi/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type POSTLoginHandler struct {
	useCase *service.LoginUserUseCase
}

func NewPOSTLoginHandler(useCase *service.LoginUserUseCase) *POSTLoginHandler {
	return &POSTLoginHandler{useCase: useCase}
}

type POSTLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type POSTLoginResponse struct {
	Token []byte `json:"token"`
}

func (p *POSTLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req POSTLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	token, err := p.useCase.Login(
		r.Context(),
		service.LoginUserCommand{
			Email:    req.Email,
			Password: []byte(req.Password),
		},
	)
	log.Print(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)

}
