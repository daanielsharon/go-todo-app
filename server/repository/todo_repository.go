package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type TodoRepository interface {
	SaveTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList) *domain.TodoList
	DeleteTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList)
	FindTodoByUsername(ctx context.Context, db *sql.DB, user *domain.User)
}
