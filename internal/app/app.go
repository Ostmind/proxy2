package app

import (
	"context"
	"log/slog"
	cl "proxynum2/internal/client"
	"proxynum2/internal/config"
	handl "proxynum2/internal/server/handler"
	srv "proxynum2/internal/server/server"
	"time"
)

type App struct {
	server *srv.Server
	logger *slog.Logger
	client *cl.ClientProxy
}

func New(logger *slog.Logger, cfg *config.AppConfig) *App {

	client := cl.New(cfg.Client, logger)

	handler := handl.New(logger, client)

	server := srv.New(logger, &cfg.Server, handler)

	return &App{
		server: server,
		logger: logger,
		client: client,
	}
}

func (a App) Run() {
	a.logger.Info("Starting app...")
	go a.server.Run()
}

func (a App) Stop(shutdownTimeout time.Duration) {
	a.logger.Info("Stopping app...")

	timeout := shutdownTimeout * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	doneCh := make(chan error)
	go func() {
		doneCh <- a.server.Stop(ctx)
	}()

	select {
	case err := <-doneCh:
		if err != nil {
			a.logger.Error("Error while stopping server: %v", err)
		}
		a.logger.Info("App has been stopped gracefully")

	case <-ctx.Done():
		a.logger.Warn("App stopped forced")
	}
}
