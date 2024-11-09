package repository

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

type ProfileRepository interface {
	GetProfileByID(ctx context.Context, userID string) (entity.Profile, error)
	GetProfileByUserID(ctx context.Context, userID string) (entity.Profile, error)
	AddProfile(ctx context.Context, profile entity.Profile) (entity.Profile, error)
	UpdateProfileByUserID(ctx context.Context, profile entity.Profile) (entity.Profile, error)
	GetRandomProfilesInSameDistrict(ctx context.Context, loggedInUserID, districtID string, limit int) ([]entity.Profile, error)
}
