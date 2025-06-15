package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	application "proxynum2/internal/app"
	"proxynum2/internal/config"
	"proxynum2/internal/server/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("No config cannot start server", slog.Any("error", err))
	}

	sloger := logger.SetupLogger()
	sloger.Info("starting proxy")

	app := application.New(sloger, cfg)

	app.Run()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	sloger.Info("Recieved interrupt signal")
	app.Stop(5)
}
