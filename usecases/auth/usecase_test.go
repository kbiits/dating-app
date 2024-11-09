package auth_usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/kbiits/dealls-take-home-test/domain/repository"
	jwt_util "github.com/kbiits/dealls-take-home-test/utils/jwt"
)

func Test_authUsecase_Login(t *testing.T) {
	type fields struct {
		txRepo      repository.TxRepo
		userRepo    repository.UserRepository
		profileRepo repository.ProfileRepository
		jwtUtil     *jwt_util.JwtUtil
	}
	type args struct {
		ctx  context.Context
		spec LoginSpec
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    LoginResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &authUsecase{
				txRepo:      tt.fields.txRepo,
				userRepo:    tt.fields.userRepo,
				profileRepo: tt.fields.profileRepo,
				jwtUtil:     tt.fields.jwtUtil,
			}
			got, err := uc.Login(tt.args.ctx, tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("authUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authUsecase.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
