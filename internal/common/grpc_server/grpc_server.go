package grpc_server

import (
	"fmt"
	"log"
	"net"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/minhvuongrbs/financial-service-example/pkg/grpc_interceptor"
	_ "github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Host string `json:"host" mapstructure:"host" yaml:"host"`
	Port int    `json:"port" mapstructure:"port" yaml:"port"`
}

type Service interface {
	RegisterService(s grpc.ServiceRegistrar)
}

func StartServer(grpcCfg Config, services ...Service) (gracefulStop func(), cerr chan error) {
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpcService := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_interceptor.ZapLoggerUnaryServerInterceptor(grpc_interceptor.DefaultShouldLog(map[string]struct{}{}))),
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			//authenticateInterceptor,
		),
		grpc.ChainStreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.MaxConcurrentStreams(1000),
		grpc.MaxRecvMsgSize(1024*1024*50), // 50MB
	)
	for _, service := range services {
		service.RegisterService(grpcService)
	}
	reflection.Register(grpcService)
	var cerrChan = make(chan error, 1)

	go func() {
		defer close(cerrChan)

		grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", grpcCfg.Host, grpcCfg.Port))
		if err != nil {
			cerrChan <- fmt.Errorf("failed to listen: %v", err)
			return
		}
		log.Println("grpc server is running on", grpcListener.Addr().String())
		defer log.Println("grpc server is stopping")

		if err := grpcService.Serve(grpcListener); err != nil {
			cerrChan <- fmt.Errorf("failed to serve: %v", err)
			return
		}
	}()

	return grpcService.GracefulStop, cerrChan
}
