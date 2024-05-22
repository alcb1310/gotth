package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Mount("/public", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	router.Get("/", handleHome)

	return router
}
