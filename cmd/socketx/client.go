package main

import (
	"fmt"
	"log/slog"

	"github.com/kissme666/socketx/internal/config"
	"github.com/kissme666/socketx/internal/socks"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client mode",
	RunE:  clientMain,
}

func clientMain(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(cfgFile)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	slog.Info("start client")
	socksServer := socks.NewServerV5(&cfg.SocksConfig)
	err = socksServer.Run()
	if err != nil {
		return fmt.Errorf("start socks5 server: %w", err)
	}
	return nil
}
