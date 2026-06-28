package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yeahChibyke/Gauntlet/internal/server"
)

func main() {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	srv := server.NewHTTPServer(":8080", logger)

	go func() {
		logger.Info(
			"starting Gauntlet server",
			"address", srv.Addr,
		)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(
				"server exited unexpectedly",
				"error", err,
			)
			os.Exit(1)
		}
	}()

	// Wait for Ctrl+C (SIGINT) or SIGTERM.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(
			"graceful shutdown failed",
			"error", err,
		)
		os.Exit(1)
	}

	logger.Info("server stopped")
}