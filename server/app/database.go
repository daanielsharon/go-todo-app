package app

import (
	"database/sql"
	"fmt"
	"server/helper"
	"time"

	_ "github.com/lib/pq"
)

func NewDatabase(name string) *sql.DB {
	dataSourceName := fmt.Sprintf("postgresql://root:root@localhost:1234/%v?sslmode=disable", name)
	db, err := sql.Open("postgres", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
