package app

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/account"
)

type LoginHandler struct {
	jwtHandler  jwtHandler
	accountRepo accountRepo
}

func NewLoginHandler(accountRepo accountRepo, jwtHandler jwtHandler) LoginHandler {
	if jwtHandler == nil {
		panic("empty jwt handler")
	}
	if accountRepo == nil {
		panic("empty account accountRepo")
	}
	return LoginHandler{
		jwtHandler:  jwtHandler,
		accountRepo: accountRepo,
	}
}

type LoginRequest struct {
	Identity string
	Password string
}

type LoginResponse struct {
	AccountId int64
	Token     string
}

type GetAccountByIdentityRequest struct {
	Identity string
	Type     account.IdentityType
}

func (h LoginHandler) Handle(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	identityType, err := account.CheckIdentityType(req.Identity)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("invalid login type: %w", err)
	}
	if !account.IsValidPassword(req.Password) {
		return LoginResponse{}, fmt.Errorf("invalid password")
	}
	a, err := h.accountRepo.GetAccountByIdentity(ctx, req.Identity, identityType)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("get account by identity failed: %w", err)
	}
	if a == nil {
		return LoginResponse{}, fmt.Errorf("invalid account info")
	}
	if err = account.ComparePassword(a.Password, req.Password); err != nil {
		return LoginResponse{}, fmt.Errorf("invalid password: %w", err)
	}
	token, err := h.jwtHandler.GenerateToken(account.TokenData{
		AccountId:   a.Id,
		UserName:    a.Username,
		Email:       a.Email,
		PhoneNumber: a.PhoneNumber,
		FullName:    a.FullName,
	})
	if err != nil {
		return LoginResponse{}, fmt.Errorf("generate jwt token failed: %w", err)
	}

	return LoginResponse{
		AccountId: a.Id,
		Token:     token,
	}, nil
}
