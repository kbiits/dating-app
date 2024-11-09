package pg_tx_repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	"github.com/kbiits/dealls-take-home-test/repositories/postgres"
)

type TxRepo struct {
	db *sqlx.DB
}

// RunInTx implements repository.TxRepo.
func (repo *TxRepo) RunInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, postgres.CtxKeyTx, tx)
	fnErr := fn(ctx)
	if fnErr != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("fn error: %w, rollback error: %v", fnErr, err)
		}

		return fnErr
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit error: %w", err)
	}

	return nil
}

func NewTxRepo(
	db *sqlx.DB,
) repository.TxRepo {
	return &TxRepo{
		db: db,
	}
}
