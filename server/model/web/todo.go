package web

type TodoCreateRequest struct {
	Name    string `json:"name"`
	UserID  int    `json:"user_id"`
	GroupID int    `json:"group_id"`
}

type TodoCreateResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	UserID  int    `json:"user_id"`
	GroupID int    `json:"group_id"`
}

type TodoDeleteRequest struct {
	ID int64 `json:"id"`
}
