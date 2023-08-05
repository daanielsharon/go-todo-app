package domain

type TodoListInsertUpdate struct {
	ID      int64
	Name    string
	GroupID int
	UserID  int
}

type TodoGroup struct {
	ID     int64
	Name   string
	UserID int
}

type TodoPriority struct {
	ID       int64
	Priority int
}

type TodoList struct {
	ID      int64
	Name    string
	UserID  int
	GroupID int
}

type Todo struct {
	ID   int64
	Name string
	// item must be accessible by index to get the nil result
	Item     []interface{}
	Priority int64
}
