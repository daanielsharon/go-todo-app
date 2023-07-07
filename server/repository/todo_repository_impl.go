package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"server/helper"
	"server/model/domain"
	"strings"
)

type TodoRepositoryImpl struct {
}

func (t *TodoRepositoryImpl) SaveTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList) *domain.TodoList {
	var lastInsertId int64

	query := "INSERT INTO todo_list(name, group_id) VALUES($1, $2)"
	err := db.QueryRowContext(ctx, query, todo.Name, todo.GroupID).Scan(&lastInsertId)

	helper.PanicIfError(err)

	todo.ID = lastInsertId
	return todo
}

func (t *TodoRepositoryImpl) DeleteTodo(ctx context.Context, db *sql.DB, todo *domain.TodoList) {
	query := "DELETE FROM todo_list WHERE id = $1"
	_, err := db.ExecContext(ctx, query, todo.ID)

	helper.PanicIfError(err)
}

func (t *TodoRepositoryImpl) FindTodoByUsername(ctx context.Context, db *sql.DB, user *domain.User) *[]domain.Todo {
	query := `
	SELECT tg.id, tg.name, ARRAY_AGG(CONCAT('{ "id": ', tl.id, ', "name": "', tl.name, '", "group_id": ', tl.group_id, '" }') ORDER BY tl.id ASC) AS item
	FROM users AS a
	JOIN todo_group as tg ON tg.user_id = a.id
	JOIN todo_list as tl ON tl.group_id = tg.id
	GROUP BY tg.id
	WHERE username = $1
	`

	rows, err := db.QueryContext(ctx, query, user.Username)
	helper.PanicIfError(err)
	defer rows.Close()

	var todo []domain.Todo

	for rows.Next() {
		// temp to store data from database
		eachTodo := domain.TodoTemp{}
		err := rows.Scan(&eachTodo.ID, &eachTodo.Name, &eachTodo.Item)
		helper.PanicIfError(err)

		// get these into slice of strings
		itemList := strings.Split(eachTodo.Item, ",")
		// hold the temp back into array

		itemTemp := []domain.TodoList{}
		for _, itemStr := range itemList {
			todoList := domain.TodoList{}
			// unmarshal each string into todolist struct
			err := json.Unmarshal([]byte(itemStr), &todoList)
			helper.PanicIfError(err)

			// put into temp array
			itemTemp = append(itemTemp, todoList)
		}

		// append this into todo
		realTodo := domain.Todo{
			ID:   eachTodo.ID,
			Name: eachTodo.Name,
			Item: itemTemp,
		}

		todo = append(todo, realTodo)
	}

	return &todo
}
