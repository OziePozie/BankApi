package bill

import (
	"BankApi/internal/domain"
	"BankApi/internal/pkg/persistence"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
)

type Repository struct {
	pool persistence.Connection
}

func NewBillRepository(pool persistence.Connection) *Repository {
	return &Repository{pool: pool}
}

//func NewBillRepository(pool *pgxpool.Pool) *Repository {
//	return &Repository{pool: pool}
//}

func (r *Repository) Save(ctx context.Context, bill *domain.Bill) error {

	query := `INSERT INTO bills (bill_uuid, account_uuid, number, sum_limit)  VALUES ($1, $2, $3, $4) `

	_, err := r.pool.Exec(ctx, query, bill.ID(), bill.UserID, bill.Balance(), 0)

	if err != nil {
		return fmt.Errorf("insert bill: %w", err)
	}

	return nil

}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetByName(ctx context.Context, name string) (domain.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) FindAllByUserWithILIKE(ctx context.Context, billName string, offset int,
	limit int, userId uuid.UUID) (*[]domain.Bill, error) {
	//query := `SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE account_uuid = $1   '%' LIMIT $3 OFFSET $4;`
	//rows, err := r.pool.Query(ctx, query, userId, billName, limit, offset)

	//     ,
	query := `SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE account_uuid = $1  AND name ILIKE $2 LIMIT $3 OFFSET $4;`
	rows, err := r.pool.Query(ctx, query, userId, "%"+billName+"%", limit, offset)
	defer rows.Close()
	log.Print(err)
	if err != nil {
		return nil, err
	}

	var bills []domain.Bill

	for rows.Next() {
		var model Model
		rows.Scan(
			&model.billId,
			&model.accId,
			&model.sum,
			&model.limit,
			&model.billName,
		)
		bills = append(bills, model.ModelToDomain())
	}

	log.Print("bills = ", bills)
	return &bills, nil

}

func (m Model) ModelToDomain() domain.Bill {
	var bill domain.Bill

	bill.SetId(m.billId)
	bill.SetName(m.billName)
	bill.SetBalance(m.sum)
	bill.SetUserID(m.accId)

	return bill
}

type Model struct {
	billId   uuid.UUID
	accId    uuid.UUID
	sum      int
	limit    int
	billName string
}

func (r *Repository) DepositAmount(ctx context.Context, billID uuid.UUID, amount int) (balance int, err error) {

	updateQuery := "UPDATE bills SET number=number+$1 WHERE bill_uuid = $2 RETURNING number;"

	row := r.pool.QueryRow(ctx, updateQuery, amount, billID)
	var newBalance int
	err = row.Scan(&newBalance)
	if err != nil {
		return -1, err
	}

	return newBalance, nil
}

func (r *Repository) GetBillByBillIDAndUserIDEquals(ctx context.Context, userID uuid.UUID, billID uuid.UUID) (*domain.Bill, error) {
	query := `SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE account_uuid = $1  AND bill_uuid = $2;`
	row := r.pool.QueryRow(ctx, query, userID, billID)
	var model Model
	err := row.Scan(
		&model.billId,
		&model.accId,
		&model.sum,
		&model.limit,
		&model.billName,
	)
	log.Print(err)
	if err != nil {
		return nil, err
	}

	bill := model.ModelToDomain()
	return &bill, nil
}

//tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
//	if err != nil {
//		return nil, err
//	}
//	defer func() {
//		if err != nil {
//			tx.Rollback(context.TODO())
//		} else {
//			tx.Commit(context.TODO())
//		}
//	}()
//
//	//bills, err := r.FindAllByUserWithILIKE(ctx, "", 0, 1, userId)
//	//
//	//bill := *bills
//	//id := bill[0].ID()
//
//	query := "SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE account_uuid = $1  LIMIT 1 FOR UPDATE"
//
//	row := tx.QueryRow(ctx, query, userId)
//
//	var model Model
//	row.Scan(
//		&model.billId,
//		&model.accId,
//		&model.sum,
//		&model.limit,
//		&model.billName,
//	)
//

//
//	if err != nil {
//		return nil, err
//	}
//
//	err = tx.Commit(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	selectQuery := "SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE bill_uuid = $1;"
//
//	row = r.pool.QueryRow(ctx, selectQuery, bill.ID())
//
//	var m Model
//	err = row.Scan(
//		&m.billId,
//		&m.accId,
//		&m.sum,
//		&m.limit,
//		&m.billName,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	bill = m.ModelToDomain()
//	log.Print(m)
//	return &bill, nil
