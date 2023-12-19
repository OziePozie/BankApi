package bill

import (
	"BankApi/internal/domain"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewBillRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

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

	//AND name ILIKE $2     "%"+billName+"%",
	query := `SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE account_uuid = $1  LIMIT $2 OFFSET $3;`
	rows, err := r.pool.Query(ctx, query, userId, limit, offset)
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

func (r *Repository) DepositAmount(ctx context.Context, userId uuid.UUID, amount int) (*domain.Bill, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	//bills, err := r.FindAllByUserWithILIKE(ctx, "", 0, 1, userId)
	//
	//bill := *bills
	//id := bill[0].ID()

	query := "SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE account_uuid = $1  LIMIT 1 FOR UPDATE"

	row := tx.QueryRow(ctx, query, userId)

	var model Model
	row.Scan(
		&model.billId,
		&model.accId,
		&model.sum,
		&model.limit,
		&model.billName,
	)

	bill := model.ModelToDomain()
	log.Print(bill)
	updateQuery := "UPDATE bills SET number=number+$1 WHERE bill_uuid = $2;"

	_, err = tx.Exec(ctx, updateQuery, amount, bill.ID())

	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	selectQuery := "SELECT bill_uuid, account_uuid, number, sum_limit, name FROM bills WHERE bill_id = $1;"

	row = r.pool.QueryRow(ctx, selectQuery, bill.ID())

	row.Scan(
		&model.billId,
		&model.accId,
		&model.sum,
		&model.limit,
		&model.billName,
	)

	bill = model.ModelToDomain()
	log.Print(bill)
	return &bill, nil
}
