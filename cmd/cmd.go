package cmd

import (
	"github.com/minhvuongrbs/financial-service-example/cmd/start_server"
	"github.com/urfave/cli/v2"
)

func AppCommandLineInterface() *cli.App {
	appCli := cli.NewApp()
	appCli.Action = start_server.StartServerAction
	appCli.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Load configuration from file path`",
			DefaultText: "./config/local.yaml",
			Value:       "./config/local.yaml",
			Required:    false,
		},
	}

	appCli.Commands = []*cli.Command{
		start_server.StartServerCmd,
	}

	return appCli
}
