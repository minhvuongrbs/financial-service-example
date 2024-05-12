package app

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/user"
)

type LoginHandler struct {
	jwtHandler jwtHandler
	userRepo   userRepo
}

func NewLoginHandler(userRepo userRepo, jwtHandler jwtHandler) LoginHandler {
	if jwtHandler == nil {
		panic("empty jwt handler")
	}
	if userRepo == nil {
		panic("empty user userRepo")
	}
	return LoginHandler{
		jwtHandler: jwtHandler,
		userRepo:   userRepo,
	}
}

type LoginRequest struct {
	Identity string
	Password string
}

type LoginResponse struct {
	UserId int64
	Token  string
}

type GetUserByUsernameIdentityRequest struct {
	Identity string
	Type     user.IdentityType
}

func (h LoginHandler) Handle(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	identityType, err := user.CheckIdentityType(req.Identity)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("invalid login type: %w", err)
	}
	if !user.IsValidPassword(req.Password) {
		return LoginResponse{}, fmt.Errorf("invalid password")
	}
	u, err := h.userRepo.GetUserByUsernameIdentity(ctx, req.Identity, identityType)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("get user by identity failed: %w", err)
	}
	if u == nil {
		return LoginResponse{}, fmt.Errorf("invalid user info")
	}
	if err = user.ComparePassword(u.Password, req.Password); err != nil {
		return LoginResponse{}, fmt.Errorf("invalid password: %w", err)
	}
	token, err := h.jwtHandler.GenerateToken(user.TokenData{
		UserId:      u.Id,
		UserName:    u.Username,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		FullName:    u.FullName,
	})
	if err != nil {
		return LoginResponse{}, fmt.Errorf("generate jwt token failed: %w", err)
	}

	return LoginResponse{
		UserId: u.Id,
		Token:  token,
	}, nil
}
