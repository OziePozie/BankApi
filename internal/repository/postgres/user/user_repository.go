package postgres

import (
	"BankApi/internal/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {

	_, err := r.pool.Exec(ctx, "INSERT INTO accounts (first_name, email, password) values ($1,$2,$3);",
		user.Name(), user.Email(), user.PasswordHash())
	defer r.pool.Close()
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil

}

func (r *UserRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	r.pool.Exec(ctx, "SELECT ")
	return nil, errNotImplemented
}
