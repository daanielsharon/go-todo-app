package setup

func (s *Setup) TruncateAll() {
	s.db.Exec("TRUNCATE users, todo_group, todo_list")
}
