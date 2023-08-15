package setup

func (s *Setup) TruncateAll() {
	s.db.Exec("TRUNCATE users, todo_group, todo_list")
}

func (s *Setup) TruncateTodo() {
	s.db.Exec("TRUNCATE todo_group, todo_list")
}
