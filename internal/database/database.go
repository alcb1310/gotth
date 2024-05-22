package database

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"github.com/alcb1310/gotth/internal/types"
)

type Service interface {
	GetAllTodos() ([]types.Todo, error)
}

type service struct {
	db *sql.DB
}

func Connect() Service {
	dbString := os.Getenv("DATABASE_URL")
	s := &service{}

	db, err := sql.Open("pgx", dbString)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err, "method", "Connect")
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		slog.Error("Failed to ping database", "error", err, "method", "Connect")
		os.Exit(1)
	}

	query := `
        CREATE TABLE IF NOT EXISTS todo (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            completed BOOLEAN NOT NULL DEFAULT false
        );
    `

	_, err = db.Exec(query)
	if err != nil {
		slog.Error("Failed to create table", "error", err, "method", "Connect")
		os.Exit(1)
	}

	s.db = db
	return s
}
