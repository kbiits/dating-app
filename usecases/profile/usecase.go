package profile_usecase

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	ctx_util "github.com/kbiits/dealls-take-home-test/utils/ctx"
	"github.com/rs/zerolog/log"
)

type ProfileUsecase interface {
	GetLoggedInProfile(ctx context.Context) (ProfileResult, error)
	CompleteProfile(ctx context.Context, spec CompleteProfileSpec) (ProfileResult, error)
}

type profileUsecase struct {
	profileRepo repository.ProfileRepository
}

func NewProfileUsecase(
	profileRepo repository.ProfileRepository,
) ProfileUsecase {
	return &profileUsecase{
		profileRepo: profileRepo,
	}
}

func (p *profileUsecase) GetLoggedInProfile(ctx context.Context) (ProfileResult, error) {
	userID, ok := ctx_util.GetUserIDFromCtx(ctx)
	if !ok {
		return ProfileResult{}, domain_errors.ErrUserNotFound
	}

	profile, err := p.profileRepo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return ProfileResult{}, err
	}

	return ProfileResult{
		ID:            profile.ID,
		UserID:        profile.UserID,
		DisplayName:   profile.DisplayName,
		Bio:           profile.Bio,
		Gender:        profile.Gender,
		ProfilePicURL: profile.ProfilePicURL,
		Status:        profile.Status,
	}, nil
}

func (p *profileUsecase) CompleteProfile(ctx context.Context, spec CompleteProfileSpec) (ProfileResult, error) {
	userID, ok := ctx_util.GetUserIDFromCtx(ctx)
	if !ok {
		return ProfileResult{}, domain_errors.ErrUserNotFound
	}

	profile, err := p.profileRepo.GetProfileByUserID(ctx, userID)
	if err != nil {
		log.Error().Str("user_id", userID).Err(err).Msg("failed to get profile by user id")
		return ProfileResult{}, err
	}

	profile.DisplayName = spec.DisplayName
	profile.Bio = spec.Bio
	profile.DateOfBirth = spec.Dob
	profile.DistrictID = spec.DistrictID
	profile.Gender = spec.Gender
	profile.ProfilePicURL = spec.ProfilePicURL

	if profile.Status == entity.ProfileStatusUnverified && profile.ShouldStatusVerified() {
		profile.Status = entity.ProfileStatusVerified
	}

	profileFresh, err := p.profileRepo.UpdateProfileByUserID(ctx, profile)
	if err != nil {
		log.Error().Err(err).Msg("failed to update profile")
		return ProfileResult{}, err
	}

	return ProfileResult{
		ID:            profileFresh.ID,
		UserID:        profileFresh.UserID,
		DisplayName:   profileFresh.DisplayName,
		Bio:           profileFresh.Bio,
		Gender:        profileFresh.Gender,
		ProfilePicURL: profileFresh.ProfilePicURL,
		Status:        profileFresh.Status,
	}, nil
}
