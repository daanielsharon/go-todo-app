package setup

import (
	"database/sql"
	"fmt"
	"server/helper"
	"time"

	_ "github.com/lib/pq"
)

func NewTestDatabase(username string, password string) *sql.DB {
	dataSourceName := fmt.Sprintf("postgresql://%v:%v@localhost:1234/todoapp_test?sslmode=disable", username, password)
	db, err := sql.Open("postgres", dataSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
