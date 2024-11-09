//go:generate mockery --all --outpkg=repo_mocks
package repository

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

type UserRepository interface {
	AddUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserLastPurchase(ctx context.Context, userID string) (*entity.UserPurchase, error)
	AddPurchaseEntry(ctx context.Context, purchase entity.UserPurchase) error
}
