package app

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/user"
	"github.com/minhvuongrbs/financial-service-example/pkg/logging"
)

type RegisterUserHandler struct {
	userRepo userRepo
}

func NewRegisterUserHandler(repo userRepo) RegisterUserHandler {
	if repo == nil {
		panic("userRepo is empty")
	}
	return RegisterUserHandler{
		userRepo: repo,
	}
}

type RegisterUserRequest struct {
	Username    string
	PhoneNumber int64
	Email       string
	BirthDay    string
	FullName    string
	Password    string
}

func (h RegisterUserHandler) Handle(ctx context.Context, req RegisterUserRequest) error {
	logging.FromContext(ctx).Infof("handle request: %v", req)
	a, err := user.NewUser(req.Username, req.PhoneNumber, req.Email, req.BirthDay, req.FullName, req.Password)
	if err != nil {
		return fmt.Errorf("invalid user info: %w", err)
	}
	a, err = h.userRepo.CreateUser(ctx, a)
	if err != nil {
		return fmt.Errorf("repository create user failed: %w", err)
	}
	return nil
}
