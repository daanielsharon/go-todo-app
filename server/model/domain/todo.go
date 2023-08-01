package domain

type TodoListInsertUpdate struct {
	ID      int64
	Name    string
	GroupID int
	UserID  int
}

type TodoGroup struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id" `
}

type TodoPriority struct {
	ID       int64
	Priority int
}

type TodoList struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	UserID  int    `json:"user_id" `
	GroupID int    `json:"group_id" `
}

type Todo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// item must be accessible by index to get the nil result
	Item     []interface{} `json:"item"`
	Priority int64         `json:"priority"`
}
