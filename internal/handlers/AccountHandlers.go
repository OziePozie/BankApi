package handlers

import (
	"BankApi/internal/models"
	"BankApi/internal/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type AccountHandler struct {
	repo *repository.AccRepoImpl
}

func New(impl *repository.AccRepoImpl) *AccountHandler {
	return &AccountHandler{repo: impl}
}

func (receiver *AccountHandler) Registration(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var accountDetails models.AccountDetails
	json.Unmarshal(body, &accountDetails)
	_, allAccounts := receiver.repo.Create(accountDetails)
	if allAccounts != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
