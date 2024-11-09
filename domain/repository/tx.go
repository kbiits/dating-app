package repository

import "context"

type TxRepo interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}
