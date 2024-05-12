package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/adapters/repository/user"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/adapters/user_token"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/app"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/ports/authgrpc"
	"github.com/minhvuongrbs/financial-service-example/internal/common/grpc_server"
	promotiongrpc "github.com/minhvuongrbs/financial-service-example/internal/promotion/ports/grpc"
)

func NewGrpcServices(conf config.Config, infra infrastructureDependencies, adapters adapters) ([]grpc_server.Service, error) {
	authService, err := NewAuthService(conf, infra, adapters)
	if err != nil {
		return nil, fmt.Errorf("failed to new auth service: %w", err)
	}

	promotionAdminService, err := NewPromotionAdminService(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to new promotion admin service: %w", err)
	}

	return []grpc_server.Service{
		authService,
		promotionAdminService,
	}, nil
}

func NewAuthService(conf config.Config, infra infrastructureDependencies, adapters adapters) (authgrpc.Service, error) {
	userRepo := user.NewRepository(infra.db)
	jwtHandler, err := user_token.NewJWTToken(conf.JwtToken)
	if err != nil {
		return authgrpc.Service{}, fmt.Errorf("new jwt handler failed: %w", err)
	}

	authApplication := authgrpc.Application{
		RegisterUser: app.NewRegisterUserHandler(userRepo),
		LoginUser:    app.NewLoginHandler(userRepo, jwtHandler),
	}
	authService := authgrpc.NewAuthService(authApplication)
	return authService, nil
}

func NewPromotionAdminService(conf config.Config) (promotiongrpc.AdminService, error) {
	service, err := promotiongrpc.NewPromotionAdminService(conf)
	if err != nil {
		return promotiongrpc.AdminService{}, fmt.Errorf("failed to init promotion admin service: %w", err)
	}
	return service, nil
}
