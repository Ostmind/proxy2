package app

import (
	"context"
	"log/slog"
	"net/http"
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

	client := cl.New(cfg.Client, logger, &http.Client{Timeout: cfg.Client.Timeout})

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

func (a App) Stop(ctx context.Context, shutdownTimeout time.Duration) {
	a.logger.Info("Stopping app...")

	timeout := shutdownTimeout * time.Second
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	doneCh := make(chan error)
	go func() {
		doneCh <- a.server.Stop(ctxWithTimeout)
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
