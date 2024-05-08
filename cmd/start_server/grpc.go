package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	grpc_server "github.com/minhvuongrbs/financial-service-example/internal/ports/grpc"
	"github.com/minhvuongrbs/financial-service-example/internal/ports/grpc/auth"
)

func NewGrpcServices(cfg config.Config, infra infrastructureDependencies, adapters *Adapters) ([]grpc_server.Service, error) {
	authService, err := NewAuthService(cfg, infra)
	if err != nil {
		return nil, fmt.Errorf("failed to new auth service: %w", err)
	}
	return []grpc_server.Service{
		authService,
	}, nil
}

func NewAuthService(_ config.Config, _ infrastructureDependencies) (auth.Service, error) {
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
	authService := auth.NewAuthService()
	return authService, nil
}
