package auth

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/api/grpc/auth"
	"github.com/minhvuongrbs/financial-service-example/internal/app/auth/login"
	"github.com/minhvuongrbs/financial-service-example/internal/app/auth/register"
	"google.golang.org/grpc"
)

type Service struct {
	auth.UnimplementedAuthServer
	App Application
}

type Application struct {
	registerAccount register.Handler
	loginAccount    login.Handler
}

func NewAuthService() Service {
	return Service{}
}

func (s Service) RegisterService(grpcServiceRegistrar grpc.ServiceRegistrar) {
	auth.RegisterAuthServer(grpcServiceRegistrar, s)
}

func (s Service) RegisterAccount(ctx context.Context,
	req *auth.RegisterAccountRequest) (*auth.RegisterAccountReply, error) {
	fmt.Printf("request: %v", req)
	err := s.App.registerAccount.Handle()
	return &auth.RegisterAccountReply{}, err
}

func (s Service) Login(context.Context, *auth.LoginRequest) (*auth.LoginReply, error) {
	err := s.App.loginAccount.Handle()
	return &auth.LoginReply{}, err
}
