package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CtxKey string

var (
	CtxKeyTx CtxKey = "pg_tx"
)

type BaseRepo struct {
	sqlx.ExtContext
}

func NewBaseRepo(db *sqlx.DB) *BaseRepo {
	return &BaseRepo{
		ExtContext: db,
	}
}

var _ sqlx.ExtContext = &BaseRepo{}

// ExecContext implements sqlx.ExtContext.
func (b *BaseRepo) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return b.ExtContext.ExecContext(ctx, query, args...)
	}

	return tx.ExecContext(ctx, query, args...)
}

// QueryContext implements sqlx.ExtContext.
func (b *BaseRepo) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return b.ExtContext.QueryContext(ctx, query, args...)
	}

	return tx.QueryContext(ctx, query, args...)
}

// QueryRowxContext implements sqlx.ExtContext.
func (b *BaseRepo) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return b.ExtContext.QueryRowxContext(ctx, query, args...)
	}

	return tx.QueryRowxContext(ctx, query, args...)
}

// QueryxContext implements sqlx.ExtContext.
func (b *BaseRepo) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return b.ExtContext.QueryxContext(ctx, query, args...)
	}

	return tx.QueryxContext(ctx, query, args...)
}

func getTxFromCtx(ctx context.Context) (*sqlx.Tx, error) {
	tx, ok := ctx.Value(CtxKeyTx).(*sqlx.Tx)
	if !ok {
		return nil, fmt.Errorf("tx not found in context")
	}

	return tx, nil
}
