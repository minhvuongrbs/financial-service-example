package start_server

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/ports/http"
	"github.com/minhvuongrbs/financial-service-example/internal/common/http_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHTTPGatewayServices(cfg config.Config) ([]http_server.GrpcGatewayServices, error) {
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
	authHttpGwService := http.NewAuthGatewayService(grpcServerConn)

	return []http_server.GrpcGatewayServices{
		authHttpGwService,
	}, nil
}
