package promotionhttp

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minhvuongrbs/financial-service-example/api/grpc/promotion"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PromotionService struct {
	grpcServerConn *grpc.ClientConn
}

func NewPromotionGatewayService(grpcClientConfig GrpcClientConfig) (*PromotionService, error) {
	grpcServerAddr := fmt.Sprintf("%s:%d", grpcClientConfig.Host, grpcClientConfig.Port)
	grpcServerConn, err := grpc.Dial(
		grpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server(%s): %w", grpcServerAddr, err)
	}
	return &PromotionService{
		grpcServerConn: grpcServerConn,
	}, nil
}

func (s *PromotionService) HTTPGatewayRegister(mux *runtime.ServeMux) error {
	err := promotion.RegisterPromotionHandler(context.Background(), mux, s.grpcServerConn)
	if err != nil {
		return fmt.Errorf("failed to register http gateway for promotion service: %w", err)
	}
	return nil
}
