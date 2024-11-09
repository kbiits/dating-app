package user_pg_repo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
)

type userRepository struct {
	db sqlx.ExtContext
}

func NewUserRepository(
	db sqlx.ExtContext,
) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	const queryTemplate = `
		INSERT INTO users (id, email, password, created_at, updated_at)
		VALUES (UUID_GENERATE_V4(), :email, :password, NOW(), NOW())
		RETURNING *
	`

	query, args, err := repo.db.BindNamed(queryTemplate, user)
	if err != nil {
		return entity.User{}, err
	}

	query = repo.db.Rebind(query)

	err = sqlx.GetContext(ctx, repo.db, &user, query, args...)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (repo *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	const query = `
		SELECT * FROM users WHERE email = $1
	`

	var user entity.User
	err := sqlx.GetContext(ctx, repo.db, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, domain_errors.ErrUserNotFound
		}

		return entity.User{}, err
	}

	return user, nil
}

func (repo *userRepository) GetUserLastPurchase(ctx context.Context, userID string) (*entity.UserPurchase, error) {
	const query = `
		SELECT * FROM user_purchases
		WHERE user_id = $1 ORDER BY purchase_date DESC LIMIT 1
	`

	var purchase entity.UserPurchase
	err := sqlx.GetContext(ctx, repo.db, &purchase, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &purchase, nil
}
