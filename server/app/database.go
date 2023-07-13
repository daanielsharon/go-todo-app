package app

import (
	"database/sql"
	"server/helper"
	"time"

	_ "github.com/lib/pq"
)

func NewDatabase() *sql.DB {
	dataSourceName := "postgresql://root:root@localhost:1234/todoapp?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
