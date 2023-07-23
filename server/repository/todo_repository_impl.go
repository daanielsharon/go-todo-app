package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"server/model/domain"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (r *TodoRepositoryImpl) InitTodoGroup(ctx context.Context, db *sql.DB, userId int) error {
	query := `
		INSERT INTO todo_group(name, user_id, priority)
		VALUES
			('todo', $1, 1),
			('in progress', $1, 2),
			('done', $1, 3)
	`

	_, err := db.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepositoryImpl) SaveTodo(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) (*domain.TodoListInsertUpdate, error) {
	var lastInsertId int64

	query := "INSERT INTO todo_list(name, group_id, user_id) VALUES($1, $2, $3) RETURNING id"
	err := db.QueryRowContext(ctx, query, todo.Name, todo.GroupID, todo.UserID).Scan(&lastInsertId)

	if err != nil {
		return &domain.TodoListInsertUpdate{}, err
	}

	todo.ID = lastInsertId
	return todo, nil
}

func (r *TodoRepositoryImpl) DeleteTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList) error {
	query := "DELETE FROM todo_list WHERE id = $1"
	_, err := db.ExecContext(ctx, query, todo.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepositoryImpl) FindTodoById(ctx context.Context, db *sql.DB, id int) (int, error) {
	query := `SELECT id FROM todo_list WHERE id = $1`

	var reqId int
	err := db.QueryRowContext(ctx, query, id).Scan(&reqId)

	if err != nil {
		return 0, err
	}

	return reqId, nil
}

func (r *TodoRepositoryImpl) FindTodoByUsername(ctx context.Context, db *sql.DB, user *domain.User) (*[]domain.Todo, error) {
	query := `
	SELECT tg.id, tg.name,
    json_agg(
        CASE
            WHEN tl.id IS NULL THEN NULL
            ELSE json_build_object('id', tl.id, 'name', tl.name)
        END
    ORDER BY tl.created_at ASC
    ) AS item,
    tg.priority
	FROM users as u
	JOIN todo_group AS tg ON tg.user_id = u.id
	LEFT JOIN todo_list AS tl ON tl.user_id = u.id AND tl.group_id = tg.id
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
		var eachTodo domain.Todo
		var itemBytes []byte
		// item is a json file. hence, we need to contain it in a separate variable
		err := rows.Scan(&eachTodo.ID, &eachTodo.Name, &itemBytes, &eachTodo.Priority)

		if err != nil {
			return &[]domain.Todo{}, err
		}

		var item []interface{}
		err = json.Unmarshal(itemBytes, &item)
		if err != nil {
			return &[]domain.Todo{}, err
		}

		eachTodo.Item = item

		todo = append(todo, eachTodo)
	}

	return &todo, nil
}
