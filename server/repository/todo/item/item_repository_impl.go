package itemrepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"server/helper"
	"server/model/domain"
)

type ItemRepositoryImpl struct {
}

func NewItemRepository() ItemRepository {
	return &ItemRepositoryImpl{}
}

func (r *ItemRepositoryImpl) Save(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) (*domain.TodoListInsertUpdate, error) {
	var lastInsertId int64

	query := "INSERT INTO todo_list(name, group_id, user_id) VALUES($1, $2, $3) RETURNING id"
	err := db.QueryRowContext(ctx, query, todo.Name, todo.GroupID, todo.UserID).Scan(&lastInsertId)
	helper.PanicIfError(err)

	todo.ID = lastInsertId
	return todo, nil
}

func (r *ItemRepositoryImpl) Update(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) *domain.TodoListInsertUpdate {
	query := "UPDATE todo_list SET group_id = $1 WHERE id = $2 AND name = $3"
	_, err := db.ExecContext(ctx, query, todo.GroupID, todo.ID, todo.Name)
	helper.PanicIfError(err)

	return todo
}

func (r *ItemRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id int64) (*domain.TodoList, error) {
	query := `SELECT id, name, group_id, user_id FROM todo_list WHERE id = $1`
	row, err := db.QueryContext(ctx, query, id)
	helper.PanicIfError(err)

	todo := &domain.TodoList{}
	if row.Next() {
		err := row.Scan(&todo.ID, &todo.Name, &todo.GroupID, &todo.UserID)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("Todo id is not found")
	}
}

func (r *ItemRepositoryImpl) FindByName(ctx context.Context, db *sql.DB, name string) (*domain.TodoList, error) {
	query := `SELECT id, name, group_id, user_id FROM todo_list WHERE name = $1`
	row, err := db.QueryContext(ctx, query, name)
	helper.PanicIfError(err)

	todo := &domain.TodoList{}
	if row.Next() {
		err := row.Scan(&todo.ID, &todo.Name, &todo.GroupID, &todo.UserID)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("Todo name is not found")
	}
}

func (r *ItemRepositoryImpl) FindByUsername(ctx context.Context, db *sql.DB, user *domain.User) (*[]domain.Todo, error) {
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
	ORDER BY tg.priority
	`

	rows, err := db.QueryContext(ctx, query, user.Username)
	helper.PanicIfError(err)

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

func (r *ItemRepositoryImpl) Delete(ctx context.Context, db *sql.DB, todo *domain.TodoList) error {
	query := "DELETE FROM todo_list WHERE id = $1"
	_, err := db.ExecContext(ctx, query, todo.ID)
	helper.PanicIfError(err)

	return nil
}
