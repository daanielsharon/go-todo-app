package domain

type TodoListInsertUpdate struct {
	ID      int64
	Name    string
	GroupID int
	UserID  int
}

type TodoList struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Item     []TodoList `json:"item"`
	Priority int64      `json:"priority"`
}

type TodoTemp struct {
	ID   int64
	Name string
	Item string
}
