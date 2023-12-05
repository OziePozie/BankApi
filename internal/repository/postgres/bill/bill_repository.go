package bill

import (
	"BankApi/internal/domain"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewBillRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) Save(ctx context.Context, bill *domain.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetAll(ctx context.Context) ([]domain.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetByName(ctx context.Context, name string) (domain.Bill, error) {
	//TODO implement me
	panic("implement me")
}
