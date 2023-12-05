package repository

import (
	"BankApi/OLD/internal/storage"
)

type Repositories struct {
	AccRepos   *AccRepoImpl
	BillsRepos *BillRepoImpl
	CardRepos  *CardRepoImpl
}

func New() *Repositories {
	storage, _ := storage.New()
	return &Repositories{
		AccRepos: &AccRepoImpl{
			s: storage},
		BillsRepos: &BillRepoImpl{
			s: storage},
		CardRepos: &CardRepoImpl{
			s: storage},
	}
}
