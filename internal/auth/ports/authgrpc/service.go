package authgrpc

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/api/grpc/auth"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/app"
	"github.com/minhvuongrbs/financial-service-example/internal/common/httpresp"
	"github.com/minhvuongrbs/financial-service-example/pkg/logging"
	"google.golang.org/grpc"
)

type Service struct {
	auth.UnimplementedAuthServer
	App Application
}

type Application struct {
	RegisterUser app.RegisterUserHandler
	LoginUser    app.LoginHandler
}

func NewAuthService(app Application) Service {
	return Service{
		App: app,
	}
}

func (s Service) RegisterService(grpcServiceRegistrar grpc.ServiceRegistrar) {
	auth.RegisterAuthServer(grpcServiceRegistrar, s)
}

func (s Service) RegisterUser(ctx context.Context,
	req *auth.RegisterUserRequest) (*auth.RegisterUserReply, error) {
	err := s.App.RegisterUser.Handle(ctx, app.RegisterUserRequest{
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
	return &auth.RegisterUserReply{
		Code:    httpresp.CodeSuccess,
		Message: httpresp.MessageSuccess,
	}, nil
}

func (s Service) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginReply, error) {
	logger := logging.FromContext(ctx).With("identity", req.Identity)
	resp, err := s.App.LoginUser.Handle(ctx, app.LoginRequest{
		Identity: req.Identity,
		Password: req.Password,
	})
	if err != nil {
		logger.Error("handle got err: %w", err)
		return nil, err
	}
	return &auth.LoginReply{
		Code:    httpresp.CodeSuccess,
		Message: httpresp.MessageSuccess,
		Data: &auth.LoginReply_Data{
			UserId: resp.UserId,
			Token:  resp.Token,
		},
	}, nil
}
