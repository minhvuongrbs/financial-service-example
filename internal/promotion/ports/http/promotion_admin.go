package promotionhttp

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minhvuongrbs/financial-service-example/api/grpc/promotion"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AdminService struct {
	grpcServerConn *grpc.ClientConn
}

type GrpcClientConfig struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

func NewPromotionAdminGatewayService(cfg GrpcClientConfig) (*AdminService, error) {
	grpcServerAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	grpcServerConn, err := grpc.Dial(grpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server(%s): %w", grpcServerAddr, err)
	}
	return &AdminService{
		grpcServerConn: grpcServerConn,
	}, nil
}

func (s *AdminService) HTTPGatewayRegister(mux *runtime.ServeMux) error {
	err := promotion.RegisterPromotionAdminHandler(context.Background(), mux, s.grpcServerConn)
	if err != nil {
		return fmt.Errorf("failed to register http gateway for promotion admin service: %w", err)
	}
	return nil
}
