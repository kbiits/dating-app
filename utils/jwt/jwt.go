package jwt_util

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kbiits/dealls-take-home-test/config"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

const (
	// because we only have 1 type of user, we can hardcode the audience and issuer
	issuer = "dating-app"
	aud    = "dating-app-users"
)

type JwtUtil struct {
	jwtConfig config.JwtConfig
}

func NewJwtUtil(
	jwtConfig config.JwtConfig,
) *JwtUtil {
	return &JwtUtil{
		jwtConfig: jwtConfig,
	}
}

func (jwtUtil *JwtUtil) GenerateToken(ctx context.Context, user entity.User) (string, error) {
	claims := newClaimsFromUser(user, time.Second*time.Duration(jwtUtil.jwtConfig.ExpirationDuration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtUtil.jwtConfig.Secret))
}

func (jwtUtil *JwtUtil) VerifyToken(ctx context.Context, token string) (*jwt.Token, bool) {
	parser := jwt.NewParser(
		jwt.WithAudience(aud),
		jwt.WithIssuer(issuer),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
	jwtParsed, err := parser.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtUtil.jwtConfig.Secret), nil
	})

	return jwtParsed, err == nil
}

func newClaimsFromUser(user entity.User, jwtDuration time.Duration) jwt.RegisteredClaims {
	jwtID := uuid.NewString()
	return jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   user.ID,
		Audience:  jwt.ClaimStrings{aud},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtDuration)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        jwtID,
	}
}
