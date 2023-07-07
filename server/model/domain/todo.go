package domain

type TodoList struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	GroupID int    `json:"group_id"`
}

type Todo struct {
	ID   int64      `json:"id"`
	Name string     `json:"name"`
	Item []TodoList `json:"item"`
}

type TodoTemp struct {
	ID   int64
	Name string
	Item string
}
