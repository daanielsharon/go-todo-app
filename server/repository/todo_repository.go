package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type TodoRepository interface {
	SaveTodo(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) (*domain.TodoListInsertUpdate, error)
	DeleteTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList) error
	FindTodoByUsername(ctx context.Context, db *sql.DB, user *domain.User) (*[]domain.Todo, error)
	FindTodoById(ctx context.Context, db *sql.DB, id int) (int, error)
}
