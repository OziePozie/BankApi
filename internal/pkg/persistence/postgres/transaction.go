package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolTransactionManager struct {
	connection *pgxpool.Pool
}

func NewPoolTransactionManager(connection *pgxpool.Pool) *PoolTransactionManager {
	return &PoolTransactionManager{
		connection: connection}
}

type transaction struct {
}

func (p PoolTransactionManager) Do(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := p.connection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	txContext := context.WithValue(ctx, transaction{}, tx)
	if err = f(txContext); err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return err
		}
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return err
		}
		return err
	}
	return err
}

type PoolConnection struct {
	pool *pgxpool.Pool
}

func NewPoolConnection(pool *pgxpool.Pool) *PoolConnection {
	return &PoolConnection{pool: pool}
}

func (c *PoolConnection) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	if tx, ok := ctx.Value(transaction{}).(pgx.Tx); ok {
		return tx.Exec(ctx, sql, arguments...)
	}
	return c.pool.Exec(ctx, sql, arguments...)
}

func (c *PoolConnection) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := ctx.Value(transaction{}).(pgx.Tx); ok {
		return tx.Query(ctx, sql, args)
	}

	return c.pool.Query(ctx, sql, args...)
}

func (c *PoolConnection) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := ctx.Value(transaction{}).(pgx.Tx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}

	return c.pool.QueryRow(ctx, sql, args...)
}
func (c *PoolConnection) Pool() *pgxpool.Pool {
	return c.pool
}
