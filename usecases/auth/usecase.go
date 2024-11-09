package auth_usecase

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	domain_errors "github.com/kbiits/dealls-take-home-test/domain/errors"
	"github.com/kbiits/dealls-take-home-test/domain/repository"
	jwt_util "github.com/kbiits/dealls-take-home-test/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(ctx context.Context, spec LoginSpec) (LoginResult, error)
	SignUp(ctx context.Context, spec SignUpSpec) (SignUpResult, error)
}

type authUsecase struct {
	txRepo      repository.TxRepo
	userRepo    repository.UserRepository
	profileRepo repository.ProfileRepository
	jwtUtil     *jwt_util.JwtUtil
}

func NewAuthUsecase(
	txRepo repository.TxRepo,
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
	jwtUtil *jwt_util.JwtUtil,
) AuthUsecase {
	return &authUsecase{
		txRepo:      txRepo,
		userRepo:    userRepo,
		profileRepo: profileRepo,
		jwtUtil:     jwtUtil,
	}
}

func (uc *authUsecase) Login(ctx context.Context, spec LoginSpec) (LoginResult, error) {
	user, err := uc.userRepo.GetUserByEmail(ctx, spec.Email)
	if err != nil {
		return LoginResult{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(spec.Password)); err != nil {
		return LoginResult{}, err
	}

	token, err := uc.jwtUtil.GenerateToken(ctx, user)
	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		UserID:      user.ID,
		AccessToken: token,
	}, nil
}

func (uc *authUsecase) SignUp(ctx context.Context, spec SignUpSpec) (SignUpResult, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(spec.Password), bcrypt.DefaultCost)
	if err != nil {
		return SignUpResult{}, err
	}

	if existingUser, _ := uc.userRepo.GetUserByEmail(ctx, spec.Email); existingUser.ID != "" {
		return SignUpResult{}, domain_errors.ErrUserAlreadyExists
	}

	user := entity.User{
		Email:    spec.Email,
		Password: string(passwordHashed),
	}

	var (
		freshUser entity.User
		token     string
	)

	err = uc.txRepo.RunInTx(ctx, func(ctx context.Context) error {
		freshUser, err = uc.userRepo.AddUser(ctx, user)
		if err != nil {
			return err
		}

		_, err := uc.profileRepo.AddProfile(ctx, entity.Profile{
			UserID:      freshUser.ID,
			DisplayName: spec.Name,
			Status:      entity.ProfileStatusUnverified,
		})
		if err != nil {
			return err
		}

		token, err = uc.jwtUtil.GenerateToken(ctx, freshUser)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return SignUpResult{}, err
	}

	return SignUpResult{
		UserID:      freshUser.ID,
		AccessToken: token,
	}, nil
}
