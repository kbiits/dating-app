package swipe_pg_repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	"github.com/rs/zerolog/log"
)

var (
	logger = log.With().Str("module", "swipe_pg_repo").Logger()
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

func (s *SwipeRepository) AddSwipeEntry(ctx context.Context, swipe entity.Swipe) error {
	const query = `
		INSERT INTO swipes (id, swiper_id, swiped_id, swipe_date, swipe_type)
		VALUES (UUID_GENERATE_V4(), :swiper_id, :swiped_id, :swipe_date, :swipe_type)
	`

	_, err := sqlx.NamedExecContext(ctx, s.db, query, swipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to add swipe entry")
		return err
	}

	return nil
}
