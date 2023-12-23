package di

import (
	"BankApi/internal/domain"
	"BankApi/internal/pkg"
	"BankApi/internal/pkg/persistence"
	"BankApi/internal/pkg/persistence/postgres"
	"BankApi/internal/repository/postgres/bill"
	"BankApi/internal/repository/postgres/user"

	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type RepoContainer struct {
	databaseURL string

	postgresPool       persistence.Connection
	transactionManager pkg.TransactionManager

	userRepository domain.UserRepository
	billRepository domain.BillRepository
}

func (c *RepoContainer) TransactionManager(ctx context.Context) pkg.TransactionManager {

	if c.transactionManager == nil {
		transactionManager, err := pgxpool.New(ctx, c.DatabaseURL())
		if err != nil {
			panic(err)
		}

		if err := transactionManager.Ping(ctx); err != nil {
			panic(err)
		}

		c.transactionManager = postgres.NewPoolTransactionManager(c.Pool(ctx).Pool())
	}

	return c.transactionManager
}
func (c *RepoContainer) Pool(ctx context.Context) persistence.Connection {
	if c.postgresPool == nil {
		postgresPool, err := pgxpool.New(ctx, c.DatabaseURL())
		if err != nil {
			panic(err)
		}

		if err := postgresPool.Ping(ctx); err != nil {
			panic(err)
		}

		c.postgresPool = postgres.NewPoolConnection(postgresPool)
	}

	return c.postgresPool
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
