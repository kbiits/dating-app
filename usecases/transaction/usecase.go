package transaction_usecase

import (
	"context"
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	ctx_util "github.com/kbiits/dealls-take-home-test/utils/ctx"
	"github.com/rs/zerolog/log"
)

var (
	logger = log.With().Str("module", "transaction_usecase").Logger()
)

type TransactionUsecase interface {
	Buy(ctx context.Context, packageID string) error
}

type transactionUsecase struct {
	userRepo           repository.UserRepository
	premiumPackageRepo repository.PremiumPackageRepository
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	premiumPackageRepo repository.PremiumPackageRepository,
) TransactionUsecase {
	return &transactionUsecase{
		userRepo:           userRepo,
		premiumPackageRepo: premiumPackageRepo,
	}
}

func (uc *transactionUsecase) Buy(ctx context.Context, packageID string) error {
	userID, ok := ctx_util.GetUserIDFromCtx(ctx)
	if !ok {
		logger.Error().Msg("error getting user id from context")
		return domain_errors.ErrUserNotFound
	}

	_, err := uc.premiumPackageRepo.GetByID(ctx, packageID)
	if err != nil {
		logger.Error().Err(err).Msg("error getting package")
		return err
	}

	// mark as paid for now
	// do transaction logic here
	err = uc.userRepo.AddPurchaseEntry(ctx, entity.UserPurchase{
		UserID:       userID,
		PackageID:    packageID,
		PurchaseDate: time.Now(),
		IsActive:     true,
	})
	if err != nil {
		logger.Error().Err(err).Msg("error adding purchase entry")
		return err
	}

	return nil
}
