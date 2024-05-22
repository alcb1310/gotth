package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handleHome)

	return router
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
