package handlers

import (
	"BankApi/OLD/internal/models"
	"BankApi/OLD/internal/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	repo *repository.AccRepoImpl
}

//func New() *AccountHandler {
//	repo := repository.New(storage)!№№@
//
//	return &AccountHandler{repo: repo.AccRepos}
//}

//func (receiver *AccountHandler) Registration(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//
//	} else {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//
//	}
//}

func (receiver *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatalln(err)
		}
		var accountDetails *models.AccountDetails

		json.Unmarshal(body, &accountDetails)

		acc, err := receiver.repo.FindAccountByLogin(accountDetails.Login)

		if acc.Password == accountDetails.Password {

			http.SetCookie(w, &http.Cookie{
				Name:  "login",
				Value: acc.Login,
			})
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
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (receiver *AccountHandler) Accounts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		if r.URL.Query().Has("id") {
			receiver.findAccount(w, r)
		} else if cookie, _ := r.Cookie("login"); cookie != nil {
			receiver.findAccountByCookie(w, r)
		} else {
			receiver.findAllAccounts(w, r)
		}
	} else if r.Method == "POST" {
		receiver.createAccount(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (receiver *AccountHandler) findAllAccounts(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	_, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var accounts []models.Account

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

func (receiver *AccountHandler) findAccountByCookie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	cookie, _ := r.Cookie("login")
	acc, err := receiver.repo.FindAccountByLogin(cookie.Value)
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

func (receiver *AccountHandler) createAccount(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var accountDetails models.AccountDetails
	json.Unmarshal(body, &accountDetails)
	_, err = receiver.repo.Create(accountDetails)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode("Account already exists")

		log.Println(err)
	} else {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Created")

	}
}
