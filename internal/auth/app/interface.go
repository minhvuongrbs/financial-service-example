package app

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/account"
)

type accountRepo interface {
	CreateAccount(ctx context.Context, a *account.Account) (*account.Account, error)
	GetAccountByIdentity(ctx context.Context, identity string, identityType account.IdentityType) (*account.Account, error)
}

type jwtHandler interface {
	GenerateToken(m account.TokenData) (string, error)
}
