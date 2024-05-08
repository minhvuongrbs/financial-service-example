package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/ports/grpc"
	"github.com/minhvuongrbs/financial-service-example/internal/common/grpc_server"
)

func NewGrpcServices(cfg config.Config, infra infrastructureDependencies, adapters adapters) ([]grpc_server.Service, error) {
	authService, err := NewAuthService(cfg, infra)
	if err != nil {
		return nil, fmt.Errorf("failed to new auth service: %w", err)
	}
	return []grpc_server.Service{
		authService,
	}, nil
}

func NewAuthService(_ config.Config, _ infrastructureDependencies) (grpc.Service, error) {
	//createUserHandler, err := create_user.NewCreateUser(adapters.UserMysqlRepository)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to new create user application: %w", err)
	//}
	//
	//loginHandler, err := login.NewLogin(adapters.ValidateUserNamePasswordWithUserUseCase, adapters.GenerateUserToken)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to new login application: %w", err)
	//}
	//
	//userApplications := users.UserApplications{
	//	CreateUserHandler: createUserHandler,
	//	Login:             loginHandler,
	//}
	authService := grpc.NewAuthService()
	return authService, nil
}
