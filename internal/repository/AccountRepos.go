package repository

import (
	"BankApi/internal/models"
	"BankApi/internal/storage"
	"database/sql"
	"fmt"
	"log"
)

type AccountRepo interface {
	FindAllAccounts(accounts *[]models.Account, storage2 *storage.Storage) (bool, error)
	Create(acc models.AccountDetails, storage2 *storage.Storage) (bool, error)
	Update(account *models.Account, storage2 *storage.Storage) (bool, error)
}

type AccRepoImpl struct {
	s *storage.Storage
	AccountRepo
}

//func New(storage *storage.Storage) *AccRepoImpl {
//
//	return &AccRepoImpl{s: storage}
//}

func (a *AccRepoImpl) FindAccountByLogin(login string) (*models.Account, error) {
	var acc models.Account
	db := a.s.Get()

	query := `SELECT * FROM accounts WHERE email = $1::TEXT;`

	rows := db.QueryRow(query, login)

	//if err != nil {
	//	return nil, err
	//}

	rows.Scan(&acc.ID,
		&acc.FirstName,
		&acc.SecondName,
		&acc.Login,
		&acc.Password)

	fmt.Println(acc)

	return &acc, nil
}

func (a *AccRepoImpl) FindAllAccounts(accounts *[]models.Account) (bool, error) {
	var acc models.Account

	db := a.s.Get()
	query := `SELECT * FROM accounts`
	stmt, err := db.Prepare(query)
	if err != nil {
		return false, sql.ErrConnDone
	}
	rows, err := stmt.Query()
	for rows.Next() {

		rows.Scan(&acc.ID,
			&acc.FirstName,
			&acc.SecondName,
			&acc.Login,
			&acc.Password)
		*accounts = append(*accounts, acc)
	}
	defer stmt.Close()
	return true, nil
}

func (a *AccRepoImpl) Create(acc models.AccountDetails) (bool, error) {
	db := a.s.Get()
	query := "INSERT INTO accounts (first_name, second_name, email, password) values ($1,$2,$3,$4);"
	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}
	res, err := stmt.Exec(acc.FirstName, acc.SecondName, acc.Login, acc.Password)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	fmt.Println(res.RowsAffected())
	return true, nil
}
func (a *AccRepoImpl) Update(account *models.Account) (bool, error) {
	db := a.s.Get()
	query := "UPDATE accounts SET first_name = $2, second_name = $3, password = $4 WHERE account_id = $1;"
	stmt, _ := db.Prepare(query)
	exec, err := stmt.Exec(account.ID, account.FirstName, account.SecondName, account.Password)
	if err != nil {
		return false, nil
	}
	log.Print(exec.RowsAffected())
	defer stmt.Close()
	return true, nil
}

func (a *AccRepoImpl) FindAccountById(id int) (*models.Account, error) {
	var acc models.Account
	db := a.s.Get()

	query := `SELECT * FROM accounts WHERE account_id = $1;`

	rows := db.QueryRow(query, id)

	//if err != nil {
	//	return nil, err
	//}

	rows.Scan(&acc.ID,
		&acc.FirstName,
		&acc.SecondName,
		&acc.Login,
		&acc.Password)

	fmt.Println(acc)

	return &acc, nil
}
