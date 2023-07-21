package setup

import "database/sql"

func TruncateAll(db *sql.DB) {
	db.Exec("TRUNCATE users, todo_group, todo_list")
}
