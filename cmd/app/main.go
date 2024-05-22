package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/alcb1310/gotth/internal/database"
	"github.com/alcb1310/gotth/internal/server"
)

func init() {
	var level slog.Level
	l := os.Getenv("APP_ENV")
	switch l {
	case "dev":
		level = slog.LevelDebug
	case "prod":
		level = slog.LevelInfo
	default:
		level = slog.LevelError
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger := slog.New(h)

	slog.SetDefault(logger)

	slog.Debug("Logger initialized")
	slog.Info("Logger initialized")
	slog.Warn("Logger initialized")
	slog.Error("Logger initialized")
}

func main() {
	db := database.Connect()
	s := server.CreateServer(db)

	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("No PORT set. Exiting")
		os.Exit(1)
	}
	slog.Info(fmt.Sprintf("Running server on port: %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s); err != nil {
		slog.Error(fmt.Sprintf("Failed to start server: %s", err))
		os.Exit(1)
	}
}
