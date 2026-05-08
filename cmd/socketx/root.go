package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kissme666/socketx/internal/config"
	"github.com/kissme666/socketx/internal/logger"
	"github.com/spf13/cobra"
)

var cfgFile string

const version = "0.0.1"

const logo = `
  _____            _        _  __   __
 / ____|          | |      | | \ \ / /
| (___   ___   ___| | _____| |_ \ V /
 \___ \ / _ \ / __| |/ / _ \ __|/   \
 ____) | (_) | (__|   <  __/ |_/ /^\ \
|_____/ \___/ \___|_|\_\___|\__\/   \/
`

func parseLevel(s string) (slog.Level, error) {
	var l slog.Level
	err := l.UnmarshalText([]byte(s))
	return l, err
}

func printBanner() {
	fmt.Print(logo)
	fmt.Printf("  version: %s\n\n", version)
}

var rootCmd = &cobra.Command{
	Use:   "socketx",
	Short: "socketx proxy bypass GFW",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		printBanner()
		cfg, err := config.Load(cfgFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "load config: %v\n", err)
			os.Exit(1)
		}
		l := cfg.Log
		level, _ := parseLevel(l.Level)
		logger.Init(logger.Options{
			Filename:   l.Filename,
			Level:      level,
			Format:     l.Format,
			MaxSize:    l.MaxSize,
			MaxBackups: l.MaxBackups,
			MaxAge:     l.MaxAge,
			Compress:   l.Compress,
		})
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file（yaml/json）")
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
