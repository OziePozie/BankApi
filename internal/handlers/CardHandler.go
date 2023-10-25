package handlers

import (
	"BankApi/internal/repository"
	"net/http"
	"strconv"
)

type CardHandler struct {
	cardRepo *repository.CardRepoImpl
	accRepo  *repository.AccRepoImpl
	billRepo *repository.BillRepoImpl
}

func (c *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("billId") {
		id, _ := strconv.Atoi(r.URL.Query().Get("billId"))
		c.cardRepo.CreateCard(id)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Card created"))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

	}
}
