package repository

import (
	m "BankApi/OLD/internal/models"
	"BankApi/OLD/internal/storage"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BillRepoImpl struct {
	s *storage.Storage
}

// func New(storage *storage.Storage) *BillRepoImpl {
//
//		return &BillRepoImpl{s: storage}
//	}
func (b *BillRepoImpl) FindAllBillsByAccountID(id int, bills *[]m.Bill) {
	db := b.s.Get()

	query := "SELECT bill_id, number, sum_limit FROM bills WHERE account_id = ($1)"
	stmt, _ := db.Prepare(query)
	rows, err := stmt.Query(id)
	if err != nil {
		return
	}
	var bill m.Bill
	for rows.Next() {

		rows.Scan(&bill.ID,
			&bill.Number,
			&bill.Limit)
		*bills = append(*bills, bill)
		fmt.Println(bill)
	}

	//for _, b := range Bills {
	//	cards := b.FindAllCardsByBillId()
	//	//fmt.Println(cards)
	//	b.Cards = cards
	//	//fmt.Println(b.Cards)
	//	bills = append(bills, b)
	//}

}
func (b *BillRepoImpl) CreateBill(id int) bool {
	connectToDB := b.s.Get()
	query := "INSERT INTO Bills (account_id, number, sum_limit) values ($1,$2,$3);"
	stmt, _ := connectToDB.Prepare(query)

	number := randomNumberBill()

	res, _ := stmt.Exec(id, number, 0)
	defer stmt.Close()

	fmt.Println(res.RowsAffected())
	return true

}

func randomNumberBill() string {
	var number string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < 20; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	return number
}
