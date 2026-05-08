package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/lumberjack.v2"
)

type Options struct {
	Filename   string
	Level      slog.Level
	// json 或 text
	Format     string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func Init(opts Options) {
	if opts.MaxSize == 0 {
		opts.MaxSize = 100
	}
	if opts.MaxBackups == 0 {
		opts.MaxBackups = 7
	}
	if opts.MaxAge == 0 {
		opts.MaxAge = 30
	}
	if opts.Format == "" {
		opts.Format = "json"
	}

	var w io.Writer = os.Stdout
	if opts.Filename != "" {
		roller := &lumberjack.Logger{
			Filename:   opts.Filename,
			MaxSize:    opts.MaxSize,
			MaxBackups: opts.MaxBackups,
			MaxAge:     opts.MaxAge,
			Compress:   opts.Compress,
		}
		w = io.MultiWriter(os.Stdout, roller)
	}

	ho := &slog.HandlerOptions{Level: opts.Level}
	var h slog.Handler
	if opts.Format == "text" {
		h = slog.NewTextHandler(w, ho)
	} else {
		h = slog.NewJSONHandler(w, ho)
	}
	slog.SetDefault(slog.New(h))
}
