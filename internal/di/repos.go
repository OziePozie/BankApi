package di

import (
	"BankApi/internal/domain"
	"BankApi/internal/repository/postgres/bill"
	"BankApi/internal/repository/postgres/user"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type RepoContainer struct {
	databaseURL string

	postgresPool *pgxpool.Pool

	userRepository domain.UserRepository
	billRepository domain.BillRepository
}

func NewRepoContainer() *RepoContainer {
	return &RepoContainer{}
}

func (c *RepoContainer) DatabaseURL() string {
	if c.databaseURL == "" {
		c.databaseURL = os.Getenv("DATABASE_URL")
	}

	return c.databaseURL
}
func (c *RepoContainer) Pool(ctx context.Context) *pgxpool.Pool {
	if c.postgresPool == nil {
		postgresPool, err := pgxpool.New(ctx, c.DatabaseURL())
		if err != nil {
			panic(err)
		}

		if err := postgresPool.Ping(ctx); err != nil {
			panic(err)
		}

		c.postgresPool = postgresPool
	}

	return c.postgresPool
}

func (c *RepoContainer) UserRepository(ctx context.Context) domain.UserRepository {
	if c.userRepository == nil {
		return user.NewUserRepository(c.Pool(ctx))
	}
	return c.userRepository
}

func (c *RepoContainer) BillRepository(ctx context.Context) domain.BillRepository {
	if c.billRepository == nil {
		return bill.NewBillRepository(c.Pool(ctx))
	}
	return c.billRepository
}

func (c *RepoContainer) SetUserRepository(ctx context.Context, repository domain.UserRepository) {
	c.userRepository = repository
}

func (c *RepoContainer) SetBillRepository(ctx context.Context, repository domain.BillRepository) {
	c.billRepository = repository

}
