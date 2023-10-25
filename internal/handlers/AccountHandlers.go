package handlers

import (
	"BankApi/internal/models"
	"BankApi/internal/repository"
	"BankApi/internal/storage"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	repo *repository.AccRepoImpl
}

func New(storage *storage.Storage) *AccountHandler {
	repo := repository.New(storage)
	return &AccountHandler{repo: repo}
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

func (receiver *AccountHandler) Accounts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("id") {
		receiver.findAccount(w, r)
	} else {
		receiver.findAllAccounts(w, r)
	}
}

func (receiver *AccountHandler) findAllAccounts(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	_, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var accounts []models.Account
	//json.Unmarshal(body, &accountDetails)
	_, err = receiver.repo.FindAllAccounts(&accounts)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(accounts)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)

}

func (receiver *AccountHandler) findAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	acc, err := receiver.repo.FindAccountById(id)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(acc)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
