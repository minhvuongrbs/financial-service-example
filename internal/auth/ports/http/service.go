package http

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minhvuongrbs/financial-service-example/api/grpc/auth"
	"google.golang.org/grpc"
)

type Service struct {
	grpcServerConn *grpc.ClientConn
}

func NewAuthGatewayService(conn *grpc.ClientConn) *Service {
	return &Service{
		grpcServerConn: conn,
	}
}

func (s *Service) HTTPGatewayRegister(mux *runtime.ServeMux) error {
	err := auth.RegisterAuthHandler(context.Background(), mux, s.grpcServerConn)
	if err != nil {
		return fmt.Errorf("failed to register http gateway for auth service: %w", err)
	}
	return nil
}
