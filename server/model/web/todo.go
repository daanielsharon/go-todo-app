package web

type TodoCreateRequest struct {
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
}
