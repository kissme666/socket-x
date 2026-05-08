package main

import (
	"fmt"
	"log/slog"

	"github.com/kissme666/socketx/internal/config"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}
		slog.Info("启动服务", "addr", cfg.Server.Addr)
		// TODO: 启动 server
		return nil
	},
}
