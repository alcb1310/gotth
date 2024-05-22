package database

import (
	"log/slog"

	"github.com/alcb1310/gotth/internal/types"
)

func (s *service) GetAllTodos() ([]types.Todo, error) {
	var todos []types.Todo

	query := `SELECT id, title, completed FROM todo;`
	rows, err := s.db.Query(query)
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo types.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	slog.Debug("GetAllTodos", "todos", todos)
	return todos, nil
}
