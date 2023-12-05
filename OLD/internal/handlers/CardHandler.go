package handlers

import (
	repository2 "BankApi/OLD/internal/repository"
	"net/http"
	"strconv"
)

type CardHandler struct {
	cardRepo *repository2.CardRepoImpl
	accRepo  *repository2.AccRepoImpl
	billRepo *repository2.BillRepoImpl
}

func (c *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
