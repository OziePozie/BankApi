package repository

import (
	"BankApi/OLD/internal/models"
	"BankApi/OLD/internal/storage"
	"database/sql"
	"fmt"
	"log"
)

type AccRepoImpl struct {
	s *storage.Storage
}

func (a *AccRepoImpl) FindAccountByLogin(login string) (*models.Account, error) {
	var acc models.Account
	db := a.s.Get()
	query := `SELECT accounts.account_id FROM accounts WHERE email = $1::TEXT`

	row := db.QueryRow(query, login)
	row.Scan(&acc.ID)

	query = `SELECT bill_id 
FROM accounts join public.bills b on 
    accounts.account_id = b.account_id WHERE b.account_id = $1`

	rows, _ := db.Query(query, acc.ID)

	var bills []int

	for rows.Next() {
		var res int
		rows.Scan(&res)
		bills = append(bills, res)
	}
	acc.Bill = bills

	//billIdChan := make(chan int)
	//done := make(chan bool)
	//var wg sync.WaitGroup
	//go func() {
	//	rows, _ := db.Query(`SELECT bill_id FROM bills WHERE account_id = $1`, acc.ID)
	//	for rows.Next() {
	//		var id int
	//		rows.Scan(&id)
	//		billIdChan <- id
	//	}
	//	close(billIdChan)
	//}()
	//
	//for id := range billIdChan {
	//	wg.Add(1)
	//	go func(id int) {
	//		defer wg.Done()
	//		fmt.Println(id)
	//		rows, _ := db.Query(`SELECT bill_id,number,sum_limit FROM bills WHERE bill_id = $1`, id)
	//		for rows.Next() {
	//			var bill models.Bill
	//			rows.Scan(&bill.ID, &bill.Number, &bill.Limit)
	//			acc.Bill = append(acc.Bill, bill)
	//		}
	//		done <- true
	//	}(id)
	//}
	//go func() {
	//	wg.Wait()   // Ожидаем завершения всех горутин
	//	close(done) // Закрываем канал для сигнализации о завершении
	//}()
	//for range done {
	//	<-done
	//}

	//var billId []int
	//
	//rows, _ := db.Query(`SELECT bill_id FROM bills WHERE account_id = $1`, acc.ID)
	//
	//for rows.Next() {
	//	var id int
	//	rows.Scan(&id)
	//	billId = append(billId, id)
	//
	//}
	//for _, id := range billId {
	//	rows, _ = db.Query(`SELECT bill_id,number,sum_limit FROM bills WHERE bill_id = $1`, id)
	//	go func() {
	//		for rows.Next() {
	//			var bill models.Bill
	//			rows.Scan(&bill.ID, &bill.Number, &bill.Limit)
	//			acc.Bill = append(acc.Bill, bill)
	//		}
	//	}()
	//}

	//c := make(chan int)
	//go func(chan int) {
	//	r := <-c
	//	rows, _ = db.Query(`SELECT bill_id,number,sum_limit FROM bills WHERE bill_id = $1`, r)
	//	for rows.Next() {
	//		var bill models.Bill
	//		rows.Scan(&bill.ID, &bill.Number, &bill.Limit)
	//		acc.Bill = append(acc.Bill, bill)
	//	}
	//
	//}(c)
	//
	//for rows.Next() {
	//	var id int
	//	rows.Scan(&id)
	//	billId = append(billId, id)
	//	c <- id
	//
	//}
	//close(c)

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

	row := db.QueryRow(query, id)

	row.Scan(&acc.ID,
		&acc.FirstName,
		&acc.SecondName,
		&acc.Login,
		&acc.Password)

	query = `SELECT bill_id 
FROM accounts join public.bills b on 
    accounts.account_id = b.account_id WHERE b.account_id = $1`

	rows, _ := db.Query(query, acc.ID)

	var bills []int

	for rows.Next() {
		var res int
		rows.Scan(&res)
		bills = append(bills, res)
	}
	acc.Bill = bills

	fmt.Println(acc)

	return &acc, nil
}
