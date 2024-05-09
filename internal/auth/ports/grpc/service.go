package grpc

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/api/grpc/auth"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/app"
	"github.com/minhvuongrbs/financial-service-example/internal/common/utils"
	"github.com/minhvuongrbs/financial-service-example/pkg/logging"
	"google.golang.org/grpc"
)

type Service struct {
	auth.UnimplementedAuthServer
	App Application
}

type Application struct {
	RegisterAccount app.RegisterAccountHandler
	LoginAccount    app.LoginHandler
}

func NewAuthService(app Application) Service {
	return Service{
		App: app,
	}
}

func (s Service) RegisterService(grpcServiceRegistrar grpc.ServiceRegistrar) {
	auth.RegisterAuthServer(grpcServiceRegistrar, s)
}

func (s Service) RegisterAccount(ctx context.Context,
	req *auth.RegisterAccountRequest) (*auth.RegisterAccountReply, error) {
	err := s.App.RegisterAccount.Handle(ctx, app.RegisterAccountRequest{
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		BirthDay:    req.Birthday,
		FullName:    req.FullName,
		Password:    req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &auth.RegisterAccountReply{
		Code:    utils.CodeSuccess,
		Message: utils.MessageSuccess,
	}, nil
}

func (s Service) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginReply, error) {
	logger := logging.FromContext(ctx).With("identity", req.Identity)
	resp, err := s.App.LoginAccount.Handle(ctx, app.LoginRequest{
		Identity: req.Identity,
		Password: req.Password,
	})
	if err != nil {
		logger.Error("handle got err: %w", err)
		return nil, err
	}
	return &auth.LoginReply{
		Code:    utils.CodeSuccess,
		Message: utils.MessageSuccess,
		Data: &auth.LoginReply_Data{
			AccountId: resp.AccountId,
			Token:     resp.Token,
		},
	}, nil
}
