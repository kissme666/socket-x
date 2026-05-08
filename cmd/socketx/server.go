package main

import (
	"fmt"
	"log/slog"

	"github.com/kissme666/socketx/internal/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}
		slog.Info("Server started", "addr", cfg.Server.Addr)
		// TODO: 启动 server
		return nil
	},
}
