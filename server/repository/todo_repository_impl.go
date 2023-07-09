package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (t *TodoRepositoryImpl) SaveTodo(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) (*domain.TodoListInsertUpdate, error) {
	var lastInsertId int64

	query := "INSERT INTO todo_list(name, group_id, user_id) VALUES($1, $2, $3)"
	err := db.QueryRowContext(ctx, query, todo.GroupID, todo.UserID).Scan(&lastInsertId)

	if err != nil {
		return &domain.TodoListInsertUpdate{}, err
	}

	todo.ID = lastInsertId
	return todo, nil
}

func (t *TodoRepositoryImpl) DeleteTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList) error {
	query := "DELETE FROM todo_list WHERE id = $1"
	_, err := db.ExecContext(ctx, query, todo.ID)

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoRepositoryImpl) FindTodoByUsername(ctx context.Context, db *sql.DB, user *domain.User) (*[]domain.Todo, error) {
	query := `
	SELECT tg.id, tg.name, json_agg(json_build_object('id', tl.id, 'name', tl.name) ORDER BY tl.created_at ASC) AS item, tg.priority
	FROM users as u
	JOIN todo_group AS tg ON tg.user_id = u.id
	JOIN todo_list AS tl ON tl.user_id = u.id AND tl.group_id = tg.id
	WHERE username = $1
	
	GROUP BY tg.id
	`

	rows, err := db.QueryContext(ctx, query, user.Username)
	if err != nil {
		return &[]domain.Todo{}, err
	}

	defer rows.Close()

	var todo []domain.Todo

	for rows.Next() {
		eachTodo := domain.Todo{}
		err := rows.Scan(&eachTodo.ID, &eachTodo.Name, &eachTodo.Item, &eachTodo.Priority)

		if err != nil {
			return &[]domain.Todo{}, err
		}

		todo = append(todo, eachTodo)
	}

	return &todo, nil
}
