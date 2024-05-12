package app

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/user"
)

type userRepo interface {
	CreateUser(ctx context.Context, a *user.User) (*user.User, error)
	GetUserByUsernameIdentity(ctx context.Context, identity string, identityType user.IdentityType) (*user.User, error)
}

type jwtHandler interface {
	GenerateToken(m user.TokenData) (string, error)
}
