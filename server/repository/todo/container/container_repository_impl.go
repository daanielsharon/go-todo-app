package repository_todo_container

import (
	"context"
	"database/sql"
	"fmt"
	"server/helper"
	"server/model/domain"
)

type ContainerRepositoryImpl struct{}

func NewContainerRepository() ContainerRepository {
	return &ContainerRepositoryImpl{}
}

func (r *ContainerRepositoryImpl) InitGroup(ctx context.Context, db *sql.DB, userId int) error {
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

func (r *ContainerRepositoryImpl) FindGroup(ctx context.Context, db *sql.DB, todo *domain.TodoGroup) (*domain.TodoGroup, error) {
	query := `SELECT id, name, user_id FROM todo_group WHERE id = $1 AND user_id = $2`
	row, err := db.QueryContext(ctx, query, todo.ID, todo.UserID)
	helper.PanicIfError(err)

	if row.Next() {
		err := row.Scan(&todo.ID, &todo.Name, &todo.UserID)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return nil, fmt.Errorf("Todo group with id %v and user_id %v does not exist", todo.ID, todo.UserID)
	}
}

func (r *ContainerRepositoryImpl) UpdatePriority(ctx context.Context, db *sql.DB, container *domain.TodoPriority) (*domain.TodoPriority, error) {
	query := `UPDATE todo_group SET priority = $1 WHERE id = $2`
	_, err := db.ExecContext(ctx, query, container.Priority, container.ID)
	if err != nil {
		return nil, err
	}

	return container, nil
}
