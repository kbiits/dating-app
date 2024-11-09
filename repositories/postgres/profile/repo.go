package profile_pg_repo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	"github.com/rs/zerolog/log"
)

type ProfileRepository struct {
	db sqlx.ExtContext
}

func NewProfileRepository(
	db sqlx.ExtContext,
) repository.ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (repo *ProfileRepository) GetProfileByID(ctx context.Context, profileID string) (entity.Profile, error) {
	const query = `
		SELECT * FROM user_profiles WHERE id = $1
	`

	var profile entity.Profile
	if err := sqlx.GetContext(ctx, repo.db, &profile, query, profileID); err != nil {
		if err == sql.ErrNoRows {
			return entity.Profile{}, domain_errors.ErrNotFound
		}

		return entity.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepository) GetProfileByUserID(ctx context.Context, userID string) (entity.Profile, error) {
	const query = `
		SELECT * FROM user_profiles WHERE user_id = $1
	`

	var profile entity.Profile
	if err := sqlx.GetContext(ctx, repo.db, &profile, query, userID); err != nil {
		if err == sql.ErrNoRows {
			return entity.Profile{}, domain_errors.ErrNotFound
		}

		return entity.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepository) AddProfile(ctx context.Context, profile entity.Profile) (entity.Profile, error) {
	const queryTemplate = `
	INSERT INTO user_profiles (id, user_id, display_name, bio, gender, date_of_birth, profile_pic_url, district_id, status)
	VALUES (
		UUID_GENERATE_V4(), :user_id, :display_name, :bio, :gender, :date_of_birth, 
		:profile_pic_url, :district_id, :status
	)
	RETURNING *
`

	query, args, err := repo.db.BindNamed(queryTemplate, profile)
	if err != nil {
		return entity.Profile{}, err
	}

	query = repo.db.Rebind(query)

	err = sqlx.GetContext(ctx, repo.db, &profile, query, args...)
	if err != nil {
		return entity.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepository) UpdateProfileByUserID(ctx context.Context, profile entity.Profile) (entity.Profile, error) {
	const queryTemplate = `
		UPDATE user_profiles SET 
			(user_id, display_name, bio, gender, date_of_birth, profile_pic_url, district_id, status)
			= (
				:user_id, :display_name, :bio, :gender, :date_of_birth, :profile_pic_url, :district_id, :status
			)
		WHERE user_id = :user_id
		RETURNING *
	`

	query, args, err := repo.db.BindNamed(queryTemplate, profile)
	if err != nil {
		return entity.Profile{}, err
	}

	query = repo.db.Rebind(query)

	err = sqlx.GetContext(ctx, repo.db, &profile, query, args...)
	if err != nil {
		return entity.Profile{}, err
	}

	return profile, nil
}

func (repo *ProfileRepository) GetRandomProfilesInSameDistrict(ctx context.Context, loggedInUserID, districtID string, limit int) ([]entity.Profile, error) {
	const query = `
		SELECT up.*
		FROM user_profiles up
		TABLESAMPLE system_rows($1)
		LEFT JOIN
			swipes s ON s.swiper_id = $4
		WHERE
			up.status = $2 AND
			s.swipe_date = CURRENT_DATE AND
			up.district_id = $3 AND
			up.user_id != $4 AND
			s.id IS NULL
	`

	var profiles []entity.Profile
	if err := sqlx.SelectContext(
		ctx, repo.db, &profiles, query, limit, entity.ProfileStatusVerified, districtID, loggedInUserID); err != nil {
		log.Error().Err(err).Msg("failed to get random profiles in same district")
		return nil, err
	}

	return profiles, nil
}
