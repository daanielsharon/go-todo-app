package containerrepo

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

func (r *ContainerRepositoryImpl) Save(ctx context.Context, db *sql.DB, container *domain.Container) *domain.Container {
	var lastInsertId int64

	query := `INSERT INTO todo_group(name, user_id, priority) VALUES($1, $2, $3) RETURNING id`
	err := db.QueryRowContext(ctx, query, container.GroupName, container.UserId, container.Priority).Scan(&lastInsertId)
	helper.PanicIfError(err)

	container.ID = lastInsertId

	return container
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
	helper.PanicIfError(err)

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

func (r *ContainerRepositoryImpl) FindTotalContainer(ctx context.Context, db *sql.DB, container *domain.Container) *uint8 {
	query := `SELECT COUNT(*) FROM users AS u JOIN todo_group AS tg ON tg.user_id = u.id WHERE user_id = $1`
	row, err := db.QueryContext(ctx, query, container.UserId)
	helper.PanicIfError(err)

	var totalContainer uint8

	if row.Next() {
		err := row.Scan(&totalContainer)
		helper.PanicIfError(err)
		return &totalContainer
	}

	return nil
}

func (r *ContainerRepositoryImpl) UpdatePriority(ctx context.Context, db *sql.DB, container *domain.TodoPriority) *domain.TodoPriority {
	query := `UPDATE todo_group SET priority = $1 WHERE id = $2`
	_, err := db.ExecContext(ctx, query, container.Priority, container.ID)
	helper.PanicIfError(err)

	return container
}
