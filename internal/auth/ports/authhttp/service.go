package authhttp

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minhvuongrbs/financial-service-example/api/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	grpcServerConn *grpc.ClientConn
}

type GrpcClientConfig struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

func NewAuthGatewayService(conf GrpcClientConfig) (*Service, error) {
	grpcServerAddr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	grpcServerConn, err := grpc.Dial(grpcServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
		//grpc.WithChainUnaryInterceptor(),
		//grpc.WithUnaryInterceptor(),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server(%s): %w", grpcServerAddr, err)
	}
	return &Service{
		grpcServerConn: grpcServerConn,
	}, nil
}

func (s *Service) HTTPGatewayRegister(mux *runtime.ServeMux) error {
	err := auth.RegisterAuthHandler(context.Background(), mux, s.grpcServerConn)
	if err != nil {
		return fmt.Errorf("failed to register http gateway for auth service: %w", err)
	}
	return nil
}
