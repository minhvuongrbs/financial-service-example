package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/adapters/account_token"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/adapters/repository/account"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/app"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/ports/grpc"
	"github.com/minhvuongrbs/financial-service-example/internal/common/grpc_server"
)

func NewGrpcServices(cfg config.Config, infra infrastructureDependencies, adapters adapters) ([]grpc_server.Service, error) {
	authService, err := NewAuthService(cfg, infra, adapters)
	if err != nil {
		return nil, fmt.Errorf("failed to new auth service: %w", err)
	}
	return []grpc_server.Service{
		authService,
	}, nil
}

func NewAuthService(conf config.Config, infra infrastructureDependencies, adapters adapters) (grpc.Service, error) {
	accountRepo := account.NewRepository(infra.db)
	jwtHandler, err := account_token.NewJWTToken(conf.JwtToken)
	if err != nil {
		return grpc.Service{}, fmt.Errorf("new jwt handler failed: %w", err)
	}

	authApplication := grpc.Application{
		RegisterAccount: app.NewRegisterAccountHandler(accountRepo),
		LoginAccount:    app.NewLoginHandler(accountRepo, jwtHandler),
	}
	authService := grpc.NewAuthService(authApplication)
	return authService, nil
}
