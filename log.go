package main

import (
	"log/slog"
	"os"
)

var logger = slog.New(
	slog.NewTextHandler(os.Stderr, nil),
)

func fatal(err error) {
	logger.Error("fatal error", slog.Any("err", err))
	os.Exit(1)
}
