package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/ports/authhttp"
	"github.com/minhvuongrbs/financial-service-example/internal/common/http_server"
	promotionhttp "github.com/minhvuongrbs/financial-service-example/internal/promotion/ports/http"
)

func NewHTTPGatewayServices(conf config.Config) ([]http_server.GrpcGatewayServices, error) {
	// new http gateway services
	authHttpGwService, err := authhttp.NewAuthGatewayService(authhttp.GrpcClientConfig{
		Host: conf.GRPC.Host,
		Port: conf.GRPC.Port,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init auth http: %w", err)
	}
	promotionAdminHttpGwService, err := promotionhttp.NewPromotionAdminGatewayService(promotionhttp.GrpcClientConfig{
		Host: conf.GRPC.Host,
		Port: conf.GRPC.Port,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init promotion admin http: %w", err)
	}
	promotionHttpGwService, err := promotionhttp.NewPromotionGatewayService(promotionhttp.GrpcClientConfig{
		Host: conf.GRPC.Host,
		Port: conf.GRPC.Port,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init promotion http: %w", err)
	}

	return []http_server.GrpcGatewayServices{
		authHttpGwService,
		promotionAdminHttpGwService,
		promotionHttpGwService,
	}, nil
}
