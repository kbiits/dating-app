package swipe_usecase

import (
	"context"
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	swipe_service "github.com/kbiits/dealls-take-home-test/domain/services/swipe"
	ctx_util "github.com/kbiits/dealls-take-home-test/utils/ctx"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

type SwipeUsecase interface {
	GetNextProfileToSwipe(ctx context.Context) (GetNextProfileResult, error)
	ClearBloomsByDate(ctx context.Context, date string) error
}

type swipeUsecase struct {
	profileRepo    repository.ProfileRepository
	swipeRepo      repository.SwipeRepository
	swipeCacheRepo repository.SwipeCacheRepository
	txRepo         repository.TxRepo
	swipeService   swipe_service.SwipeService
}

func NewSwipeUsecase(
	profileRepo repository.ProfileRepository,
	swipeRepo repository.SwipeRepository,
	swipeCacheRepo repository.SwipeCacheRepository,
	txRepo repository.TxRepo,
	swipeService swipe_service.SwipeService,
) SwipeUsecase {
	return &swipeUsecase{
		profileRepo:    profileRepo,
		swipeRepo:      swipeRepo,
		swipeCacheRepo: swipeCacheRepo,
		txRepo:         txRepo,
		swipeService:   swipeService,
	}
}

func (s *swipeUsecase) GetNextProfileToSwipe(ctx context.Context) (GetNextProfileResult, error) {
	const limitProfiles = 10
	userID, ok := ctx_util.GetUserIDFromCtx(ctx)
	if !ok {
		return GetNextProfileResult{}, domain_errors.ErrUserNotFound
	}

	loggedInProfile, err := s.profileRepo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return GetNextProfileResult{}, err
	}

	if can, err := s.swipeService.UserCanDoSwipe(ctx, loggedInProfile); err != nil {
		return GetNextProfileResult{}, err
	} else if !can {
		return GetNextProfileResult{}, domain_errors.ErrCannotDoSwipe
	}

	profiles, err := s.profileRepo.GetRandomProfilesInSameDistrict(ctx, userID, loggedInProfile.DistrictID.OrEmpty(), limitProfiles)
	if err != nil {
		log.Error().Err(err).Msg("failed to get random profiles in same district")
		return GetNextProfileResult{}, err
	}

	if len(profiles) == 0 {
		return GetNextProfileResult{}, domain_errors.ErrProfileNotFound
	}

	profileIds := lo.Map(profiles, func(profile entity.Profile, _ int) string {
		return profile.ID
	})

	date := time.Now().Format("2006-01-02")

	unswiped, err := s.swipeCacheRepo.GetNonSwiped(ctx, date, loggedInProfile.ID, profileIds)
	if err != nil {
		log.Error().Err(err).Msg("failed to get non-swiped profiles")
		return GetNextProfileResult{}, err
	}

	if len(unswiped) == 0 {
		return GetNextProfileResult{}, domain_errors.ErrProfileNotFound
	}

	var (
		nextProfileID = unswiped[0]
		nextProfile   entity.Profile
	)

	nextProfile, err = s.profileRepo.GetProfileByID(ctx, nextProfileID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get profile by id")
		return GetNextProfileResult{}, err
	}

	return GetNextProfileResult{
		DisplayName:   nextProfile.DisplayName,
		Bio:           nextProfile.Bio,
		DateOfBirth:   nextProfile.DateOfBirth,
		Gender:        nextProfile.Gender,
		ProfilePicURL: nextProfile.ProfilePicURL,
	}, nil
}

func (s *swipeUsecase) ClearBloomsByDate(ctx context.Context, date string) error {
	if err := s.swipeCacheRepo.ClearYesterdayBloomFilter(ctx, date); err != nil {
		log.Error().Err(err).Msg("failed to clear yesterday bloom filter")
		return err
	}

	return nil
}
