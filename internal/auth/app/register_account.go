package app

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/account"
)

type RegisterAccountHandler struct {
	accountRepo accountRepo
}

func NewRegisterAccountHandler(repo accountRepo) RegisterAccountHandler {
	if repo == nil {
		panic("account accountRepo is empty")
	}
	return RegisterAccountHandler{
		accountRepo: repo,
	}
}

type RegisterAccountRequest struct {
	Username    string
	PhoneNumber int64
	Email       string
	BirthDay    string
	FullName    string
	Password    string
}

func (h RegisterAccountHandler) Handle(ctx context.Context, req RegisterAccountRequest) error {
	a, err := account.NewAccount(req.Username, req.PhoneNumber, req.Email, req.BirthDay, req.FullName, req.Password)
	if err != nil {
		return fmt.Errorf("invalid account info: %w", err)
	}
	a, err = h.accountRepo.CreateAccount(ctx, a)
	if err != nil {
		return fmt.Errorf("repository create account failed: %w", err)
	}
	return nil
}
