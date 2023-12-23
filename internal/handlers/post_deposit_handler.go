package handlers

import (
	"BankApi/internal/handlers/middleware"
	"BankApi/internal/service"
	"encoding/json"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
)

type POSTDepositHandler struct {
	useCase *service.DepositUseCase
}

func NewPOSTDepositHandler(useCase *service.DepositUseCase) *POSTDepositHandler {
	return &POSTDepositHandler{useCase: useCase}
}

type POSTDepositRequest struct {
	BillID uuid.UUID `json:"billID"`
	Amount int       `json:"amount"`
}

type POSTDepositResponse struct {
	AccountBalance int `json:"accountBalance"`
}

func (handler POSTDepositHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userID := middleware.UserIDFromContext(request.Context())

	var body POSTDepositRequest
	err := json.NewDecoder(request.Body).Decode(&body)
	log.Print(err)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	command := service.DepositCommand{
		UserID: userID,
		BillID: body.BillID,
		Amount: body.Amount,
	}

	accountBalance, err := handler.useCase.Handle(request.Context(), command)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	response := POSTDepositResponse{AccountBalance: accountBalance}

	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(response)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

}
