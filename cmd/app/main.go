package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/alcb1310/gotth/internal/server"
)

func main() {
	s := server.CreateServer()

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
