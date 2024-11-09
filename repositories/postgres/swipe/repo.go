package swipe_pg_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	"github.com/rs/zerolog/log"
)

type SwipeRepository struct {
	db sqlx.ExtContext
}

func NewSwipeRepository(
	db sqlx.ExtContext,
) repository.SwipeRepository {
	return &SwipeRepository{
		db: db,
	}
}

func (s *SwipeRepository) CountUserSwipeByDate(ctx context.Context, userID string, date string) (int, error) {
	const query = `
		SELECT COUNT(*) FROM swipes WHERE swipe_date = $1 AND swiper_id = $2
	`

	var count int
	err := sqlx.GetContext(ctx, s.db, &count, query, date, userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to count user swipe by date")
		return 0, err
	}

	return count, nil
}
