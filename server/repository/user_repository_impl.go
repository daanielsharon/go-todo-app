package repository

import (
	"context"
	"database/sql"
	"errors"
	"server/helper"
	"server/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Find(ctx context.Context, db *sql.DB, username *domain.User) (*domain.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row, err := db.QueryContext(ctx, query, username.Username)

	if err != nil {
		return &domain.User{}, err
	}

	user := domain.User{}
	if row.Next() {
		err := row.Scan(&user.ID, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, db *sql.DB, user *domain.User) (*domain.User, error) {
	var userId int64
	query := "INSERT INTO users(username, password) VALUES($1, $2) RETURNING id"
	err := db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&userId)

	if err != nil {
		return &domain.User{}, err
	}

	user.ID = userId
	return user, nil
}
