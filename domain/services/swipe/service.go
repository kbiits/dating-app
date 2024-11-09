package swipe_service

import (
	"context"
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	"github.com/rs/zerolog/log"
)

type SwipeService interface {
	UserCanDoSwipe(ctx context.Context, profile entity.Profile) (bool, error)
}

type swipeService struct {
	userRepo           repository.UserRepository
	swipeRepo          repository.SwipeRepository
	premiumPackageRepo repository.PremiumPackageRepository
}

func NewSwipeService(
	userRepo repository.UserRepository,
	swipeRepo repository.SwipeRepository,
	premiumPackageRepo repository.PremiumPackageRepository,
) SwipeService {
	return &swipeService{
		userRepo:           userRepo,
		swipeRepo:          swipeRepo,
		premiumPackageRepo: premiumPackageRepo,
	}
}

const (
	basicUserSwipeQuota = 10
)

func (s *swipeService) UserCanDoSwipe(ctx context.Context, profile entity.Profile) (bool, error) {
	if profile.Status != entity.ProfileStatusVerified {
		log.Info().Msg("profile is not verified to do swipe")
		return false, nil
	}

	userPurchase, err := s.userRepo.GetUserLastPurchase(ctx, profile.UserID)
	if err != nil {
		return false, err
	}

	if userPurchase == nil {
		if quota, err := s.checkUserSwipeQuota(ctx, nil, profile); err != nil {
			return false, err
		} else if quota <= 0 {
			log.Debug().Msg("user has no swipe quota")
			return false, nil
		}

		return true, nil
	}

	// check purchase status and expiry
	if !userPurchase.IsActive {
		log.Debug().Msg("premium package is not active")
		return false, nil
	}

	premiumPackage, err := s.premiumPackageRepo.GetByID(ctx, userPurchase.PackageID)
	if err != nil {
		return false, err
	}

	minutesSinceBuy := time.Since(userPurchase.PurchaseDate).Minutes()
	if minutesSinceBuy > float64(premiumPackage.Validity) {
		log.Debug().Msg("premium package is expired")
		return false, nil
	}

	quota, err := s.checkUserSwipeQuota(ctx, &premiumPackage, profile)
	if err != nil {
		return false, err
	}

	if quota <= 0 {
		log.Debug().Msg("user has no swipe quota")
		return false, nil
	}

	return true, nil
}

func (s *swipeService) checkUserSwipeQuota(ctx context.Context, premiumPackage *entity.PremiumPackage, profile entity.Profile) (int, error) {
	if premiumPackage != nil && premiumPackage.Config.UnlimitedQuota {
		return 1_000, nil
	}

	date := time.Now().Format("2006-01-02")
	count, err := s.swipeRepo.CountUserSwipeByDate(ctx, profile.UserID, date)
	if err != nil {
		return 0, err
	}

	if premiumPackage == nil {
		return basicUserSwipeQuota - count, nil
	}

	return premiumPackage.Config.QuotaPerDay - count, nil
}
