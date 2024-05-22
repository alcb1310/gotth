package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alcb1310/gotth/internal/database"
)

type service struct {
	db database.Service
}

func CreateServer(db database.Service) *chi.Mux {
	s := &service{
		db: db,
	}
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Mount("/public", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	router.Get("/", s.handleHome)
	router.Get("/todos", s.handleTodos)

	return router
}
