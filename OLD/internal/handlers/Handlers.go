package handlers

import (
	"BankApi/OLD/internal/repository"
)

type Handlers struct {
	AccHandler  *AccountHandler
	BillHandler *BillHandler
	CardHandler *CardHandler
}

func New() *Handlers {
	repos := repository.New()
	return &Handlers{
		AccHandler: &AccountHandler{
			repo: repos.AccRepos,
		},
		BillHandler: &BillHandler{
			repo:    repos.BillsRepos,
			accRepo: repos.AccRepos,
		},
		CardHandler: &CardHandler{
			cardRepo: repos.CardRepos,
			accRepo:  repos.AccRepos,
			billRepo: repos.BillsRepos,
		},
	}

}
