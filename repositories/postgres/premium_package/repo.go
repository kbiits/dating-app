package premium_package_pg_repo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
)

type PremiumPackageRepository struct {
	db sqlx.ExtContext
}

func NewPremiumPackageRepo(
	db sqlx.ExtContext,
) repository.PremiumPackageRepository {
	return &PremiumPackageRepository{
		db: db,
	}
}

func (p *PremiumPackageRepository) GetByID(ctx context.Context, id string) (entity.PremiumPackage, error) {
	const query = `
		SELECT * FROM premium_packages WHERE id = $1 LIMIT 1
	`

	var premiumPackage entity.PremiumPackage
	err := sqlx.GetContext(ctx, p.db, &premiumPackage, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.PremiumPackage{}, domain_errors.ErrPremiumPackageNotFound
		}

		return entity.PremiumPackage{}, err
	}

	return premiumPackage, nil
}
