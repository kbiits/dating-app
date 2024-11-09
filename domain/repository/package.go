package repository

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

type PremiumPackageRepository interface {
	GetByID(ctx context.Context, id string) (entity.PremiumPackage, error)
}
