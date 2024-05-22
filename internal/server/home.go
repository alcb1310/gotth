package server

import (
	"net/http"

	"github.com/alcb1310/gotth/internal/views/home"
)

func (s *service) handleHome(w http.ResponseWriter, r *http.Request) {
	c := home.Index()

	c.Render(r.Context(), w)
}
