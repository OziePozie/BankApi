package handlers

import (
	"BankApi/internal/models"
	"BankApi/internal/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BillHandler struct {
	repo    *repository.BillRepoImpl
	accRepo *repository.AccRepoImpl
}

func (h *BillHandler) Bills(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		defer r.Body.Close()
		_, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatalln(err)
		}

		cookie, _ := r.Cookie("login")

		acc, err := h.accRepo.FindAccountByLogin(cookie.Value)

		var bills []models.Bill

		h.repo.FindAllBillsByAccountID(acc.ID, &bills)

		if err != nil {
			return
		}
		marshal, err := json.Marshal(bills)
		if err != nil {
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshal)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *BillHandler) CreateBill(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cookie, _ := r.Cookie("login")
		acc, _ := h.accRepo.FindAccountByLogin(cookie.Value)
		h.repo.CreateBill(acc.ID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *BillHandler) SetLimit(w http.ResponseWriter, r *http.Request) {

}
