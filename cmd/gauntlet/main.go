package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yeahChibyke/Gauntlet/internal/provider/nvidia"
	"github.com/yeahChibyke/Gauntlet/internal/server"
	"github.com/yeahChibyke/Gauntlet/internal/service"
)

func main() {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	nvidiaProvider, err := nvidia.NewProvider()
	if err != nil {
		logger.Error(
			"failed to initialize NVIDIA provider",
			"error", err,
		)
		os.Exit(1)
	}

	responseService := service.NewResponseService(
		nvidiaProvider,
	)

	srv := server.NewHTTPServer(
		":8080",
		logger,
		responseService,
	)

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
