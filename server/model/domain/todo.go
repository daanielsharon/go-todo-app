package domain

type TodoListInsertUpdate struct {
	ID      int64
	Name    string
	GroupID int
	UserID  int
}

type TodoList struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
