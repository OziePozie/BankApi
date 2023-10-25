package handlers

import (
	"BankApi/internal/models"
	"BankApi/internal/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AccountHandler struct {
	repo *repository.AccRepoImpl
}

func New() *AccountHandler {

	return &AccountHandler{repo: repository.New()}
}

func (receiver *AccountHandler) Registration(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var accountDetails models.AccountDetails
	json.Unmarshal(body, &accountDetails)
	_, err = receiver.repo.Create(accountDetails)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
func (receiver *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var accountDetails *models.AccountDetails

	json.Unmarshal(body, &accountDetails)
	fmt.Println(accountDetails.Login)
	acc, err := receiver.repo.FindAccountByLogin(accountDetails.Login)

	if acc.Password == accountDetails.Password {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Successful login")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Unsuccessful login")
	}

	if err != nil {
		return
	}

}
