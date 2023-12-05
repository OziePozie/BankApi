package repository

import (
	"BankApi/internal/storage"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type CardRepoImpl struct {
	s *storage.Storage
}

func (c *CardRepoImpl) CreateCard(billId int) bool {
	db := c.s.Get()
	query := `INSERT INTO cards (bill_id, number, cvv, expiration_date, iscardactive)
            values ($1,$2,$3, $4, $5);`

	stmt, _ := db.Prepare(query)

	cvv, number := randomCVVAndNumber()

	res, _ := stmt.Exec(billId,
		number,
		cvv,
		time.Now().AddDate(4, 0, 0),
		true,
	)
	defer stmt.Close()

	fmt.Println(res.RowsAffected())
	return true

}
func randomCVVAndNumber() (string, string) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var number string
	for i := 0; i < 16; i++ {
		number += strconv.Itoa(r.Intn(10))
	}
	var cvv string
	for i := 0; i < 3; i++ {
		cvv += strconv.Itoa(r.Intn(10))
	}
	return cvv, number
}
