package domain

type TodoListInsertUpdate struct {
	ID      int64
	Name    string
	GroupID int
	UserID  int
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

type TodoTemp struct {
	ID   int64
	Name string
	Item string
}
