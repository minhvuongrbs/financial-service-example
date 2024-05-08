package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	http_gateway "github.com/minhvuongrbs/financial-service-example/internal/ports/http"
	"github.com/minhvuongrbs/financial-service-example/internal/ports/http/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHTTPGatewayServices(cfg config.Config) ([]http_gateway.GrpcGatewayServices, error) {
	grpcServerAddr := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)
	grpcServerConn, err := grpc.Dial(grpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
		//grpc.WithChainUnaryInterceptor(),
		//grpc.WithUnaryInterceptor(),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server(%s): %w", grpcServerAddr, err)
	}

	// new http gateway services
	authHttpGwService := auth.NewAuthGatewayService(grpcServerConn)

	return []http_gateway.GrpcGatewayServices{
		authHttpGwService,
	}, nil
}
