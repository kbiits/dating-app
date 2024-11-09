package repository

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

type SwipeRepository interface {
	CountUserSwipeByDate(ctx context.Context, userID string, date string) (int, error)
	AddSwipeEntry(ctx context.Context, swipe entity.Swipe) error
}

type SwipeCacheRepository interface {
	ClearYesterdayBloomFilter(ctx context.Context, date string) error
	GetNonSwiped(ctx context.Context, date, swiperProfileID string, profileIDs []string) ([]string, error)
	MarkAsSwiped(ctx context.Context, date, swiperProfileID string, profileID string) (bool, error)
}
