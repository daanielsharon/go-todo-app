package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type TodoRepositoryImpl struct {
	db *sql.DB
}

func (t *TodoRepositoryImpl) Save(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	var lastInsertId int64

	query := "INSERT INTO todo(name, container_id) VALUES($1, $2)"
	err := t.db.QueryRowContext(ctx, query, todo.Name, todo.ContainerId).Scan(&lastInsertId)

	if err != nil {
		return &domain.Todo{}, err
	}

	todo.ID = lastInsertId
	return todo, nil
}