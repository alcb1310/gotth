package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/alcb1310/gotth/internal/views/todos"
)

func (s *service) handleTodos(w http.ResponseWriter, r *http.Request) {
	t, err := s.db.GetAllTodos()
	if err != nil {
		slog.Error("Failed to get todos", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	c := todos.ListTodo(t)

	time.Sleep(2 * time.Second)

	c.Render(r.Context(), w)
}
