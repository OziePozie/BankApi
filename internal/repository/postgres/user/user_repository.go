package user

import (
	"BankApi/internal/domain"
	"BankApi/internal/repository"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Save(ctx context.Context, user *domain.User) error {

	log.Print(user.Email())

	_, err := r.pool.Exec(ctx, "INSERT INTO accounts (first_name, email, password) values ($1,$2,$3);",
		user.Name(), user.Email(), user.PasswordHash())
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil

}

func (r *Repository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	r.pool.Exec(ctx, "SELECT ")
	return nil, repository.ErrNotImplement
}
