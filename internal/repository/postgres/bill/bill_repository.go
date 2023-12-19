package bill

import (
	"BankApi/internal/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
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
