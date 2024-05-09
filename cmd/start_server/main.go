package start_server

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/internal/common/grpc_server"
	"github.com/minhvuongrbs/financial-service-example/internal/common/http_server"
	"github.com/minhvuongrbs/financial-service-example/pkg/logging"
	"github.com/urfave/cli/v2"
)

var (
	StartServerCmd = &cli.Command{
		Name:   "server",
		Usage:  "run http server",
		Action: StartServerAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from file path`",
				DefaultText: "./config/local.yaml",
				Value:       "./config/local.yaml",
				Required:    false,
			},
		},
	}
)

func StartServerAction(cmdCLI *cli.Context) error {
	cfgPath := cmdCLI.String("config")
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to load config from path\"%s\": %w", cfgPath, err)
	}
	return StartHTTPServer(cfg)
}

func StartHTTPServer(cfg config.Config) error {
	if cfg.Env == "local" {
		bs, err := json.Marshal(cfg)
		if err != nil {
			return fmt.Errorf("failed to marshal config: %w", err)
		}
		fmt.Println("Start server with config:", string(bs))
	}

	err := logging.InitLogger(&cfg.Logger)
	if err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	l := logging.FromContext(context.Background())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	cerr := make(chan error)
	go func() {
		for _err := range cerr {
			l.Errorw("server error got error", "error", _err)
			stop <- syscall.SIGTERM
		}
	}()

	//init infrastructure
	infra, err := initInfrastructure(cfg)
	if err != nil {
		return fmt.Errorf("failed to init infrastructure: %w", err)
	}

	serviceAdapters, err := newAdapters(cfg, infra)
	if err != nil {
		return fmt.Errorf("failed to new serviceAdapters: %w", err)
	}

	// new application
	grpcServices, err := NewGrpcServices(cfg, infra, serviceAdapters)
	if err != nil {
		return fmt.Errorf("failed to new grpc services: %w", err)
	}

	// start server
	grpcStop, cgrpcerr := grpc_server.StartServer(cfg.GRPC, grpcServices...)
	go func() {
		for gerr := range cgrpcerr {
			cerr <- fmt.Errorf("grpc server error: %w", gerr)
		}
	}()
	defer grpcStop()

	// start http server
	httpgwServices, err := NewHTTPGatewayServices(cfg)
	if err != nil {
		return fmt.Errorf("failed to new http gateway services: %w", err)
	}

	httpServices := []http_server.HTTPService{
		//monitor.PprofService{},
		//monitor.PrometheusService{},
	}
	httpStop, cherr := http_server.StartServer(cfg.HTTP, httpgwServices, httpServices)
	go func() {
		for herr := range cherr {
			cerr <- fmt.Errorf("http server error: %w", herr)
		}
	}()
	defer httpStop()

	l.Infow("server started")
	<-stop
	l.Infow("server stopping")
	return nil
}
