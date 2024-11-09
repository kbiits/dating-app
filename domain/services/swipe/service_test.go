package swipe_service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
	repo_mocks "github.com/kbiits/dealls-take-home-test/domain/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_swipeService_UserCanDoSwipe(t *testing.T) {
	type fields struct {
		userRepo           *repo_mocks.UserRepository
		swipeRepo          *repo_mocks.SwipeRepository
		premiumPackageRepo *repo_mocks.PremiumPackageRepository
	}

	field := fields{
		userRepo:           repo_mocks.NewUserRepository(t),
		swipeRepo:          repo_mocks.NewSwipeRepository(t),
		premiumPackageRepo: repo_mocks.NewPremiumPackageRepository(t),
	}

	type args struct {
		ctx     context.Context
		profile entity.Profile
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		wantErr  error
		mockFunc func(f fields)
	}{
		{
			name: "user can't do swipe because profile is not verified",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					Status: entity.ProfileStatusDeactivated,
				},
			},
			mockFunc: func(f fields) {},
			want:     false,
		},
		{
			name: "err get last purchase",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(nil, fmt.Errorf("error get last purchase")).
					Once()
			},
			wantErr: fmt.Errorf("error get last purchase"),
			want:    false,
		},
		{
			name: "no purchase and get swipe quota error",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(nil, nil).
					Once()
				f.swipeRepo.On("CountUserSwipeByDate", mock.Anything, "user-id-1", time.Now().Format("2006-01-02")).
					Return(0, fmt.Errorf("err get swipe quota")).
					Once()
			},
			want:    false,
			wantErr: fmt.Errorf("err get swipe quota"),
		},
		{
			name: "no purchase and no swipe quota",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(nil, nil).
					Once()
				f.swipeRepo.On("CountUserSwipeByDate", mock.Anything, "user-id-1", time.Now().Format("2006-01-02")).
					Return(basicUserSwipeQuota, nil).
					Once()
			},
			want: false,
		},
		{
			name: "no purchase and but still has quota",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(nil, nil).
					Once()
				f.swipeRepo.On("CountUserSwipeByDate", mock.Anything, "user-id-1", time.Now().Format("2006-01-02")).
					Return(basicUserSwipeQuota-5, nil).
					Once()
			},
			want: true,
		},
		{
			name: "has purchase but not active",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(&entity.UserPurchase{
						IsActive: false,
					}, nil).
					Once()
			},
			want: false,
		},
		{
			name: "has purchase but premium package not found",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(&entity.UserPurchase{
						PackageID: "package-id-1",
						IsActive:  true,
					}, nil).
					Once()

				f.premiumPackageRepo.On("GetByID", mock.Anything, "package-id-1").
					Return(entity.PremiumPackage{}, fmt.Errorf("error get premium package")).
					Once()
			},
			want:    false,
			wantErr: fmt.Errorf("error get premium package"),
		},
		{
			name: "has purchase but already expired",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(&entity.UserPurchase{
						PackageID:    "package-id-1",
						IsActive:     true,
						PurchaseDate: time.Now().Add(-time.Hour * 1),
					}, nil).
					Once()

				f.premiumPackageRepo.On("GetByID", mock.Anything, "package-id-1").
					Return(entity.PremiumPackage{
						Validity: 10,
					}, nil).
					Once()
			},
			want:    false,
			wantErr: nil,
		},
		{
			name: "has purchase but the quota is empty",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(&entity.UserPurchase{
						PackageID:    "package-id-1",
						IsActive:     true,
						PurchaseDate: time.Now(),
					}, nil).
					Once()

				f.premiumPackageRepo.On("GetByID", mock.Anything, "package-id-1").
					Return(entity.PremiumPackage{
						Validity: 10,
						Config: entity.PremiumPackageConfig{
							QuotaPerDay: 5,
						},
					}, nil).
					Once()

				f.swipeRepo.On("CountUserSwipeByDate", mock.Anything, "user-id-1", time.Now().Format("2006-01-02")).
					Return(5, nil).
					Once()
			},
		},
		{
			name: "has purchase and unlimited quota",
			args: args{
				ctx: context.Background(),
				profile: entity.Profile{
					UserID: "user-id-1",
					Status: entity.ProfileStatusVerified,
				},
			},
			mockFunc: func(f fields) {
				f.userRepo.On("GetUserLastPurchase", mock.Anything, "user-id-1").
					Return(&entity.UserPurchase{
						PackageID:    "package-id-1",
						IsActive:     true,
						PurchaseDate: time.Now(),
					}, nil).
					Once()

				f.premiumPackageRepo.On("GetByID", mock.Anything, "package-id-1").
					Return(entity.PremiumPackage{
						Validity: 10,
						Config: entity.PremiumPackageConfig{
							UnlimitedQuota: true,
						},
					}, nil).
					Once()
			},
			want:    true,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &swipeService{
				userRepo:           field.userRepo,
				swipeRepo:          field.swipeRepo,
				premiumPackageRepo: field.premiumPackageRepo,
			}
			tt.mockFunc(field)

			got, err := s.UserCanDoSwipe(tt.args.ctx, tt.args.profile)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}

			require.EqualValues(t, tt.want, got)
		})
	}
}
