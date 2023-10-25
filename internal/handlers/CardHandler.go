package handlers

import (
	"BankApi/internal/repository"
	"net/http"
)

type CardHandler struct {
	cardRepo *repository.CardRepoImpl
	accRepo  *repository.AccRepoImpl
	billRepo *repository.BillRepoImpl
}

func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {

}
