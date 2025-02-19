package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"

	"preparation/go-projects/govault/config"
	"preparation/go-projects/govault/internal/handler"
)

func main() {
	var config config.Config

	err := envconfig.Process("", &config)
	if err != nil {
		slog.Error("env file is missing")
	}

	httpServer := &http.Server{Addr: ":" + config.APP.PORT}

	ctx := context.Background()
	slog.Log(ctx, slog.LevelInfo, "GoVault is running on", "port", config.APP.PORT)

	http.HandleFunc("/set", handler.SetKey)
	http.HandleFunc("/get", handler.GetKey)
	http.HandleFunc("/delete", handler.DeleteKey)
	http.HandleFunc("/set_with_ttl", handler.SetKeyTTL)
	http.HandleFunc("/select_db", handler.SelectDB)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			slog.Error("Server error", "error", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down server", "error", err)
	} else {
		slog.Info("Server shut down gracefully")
	}

}
