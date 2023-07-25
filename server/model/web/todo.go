package web

type TodoCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	UserID  int    `json:"user_id" validate:"required"`
	GroupID int    `json:"group_id" validate:"required"`
}

type TodoCreateUpdateResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	UserID  int    `json:"user_id"`
	GroupID int    `json:"group_id"`
}

type TodoGetRequest struct {
	Username string `json:"username" validate:"required,min=1,max=30"`
}

type TodoGetResponse struct {
	ID        int64       `json:"id"`
	GroupName string      `json:"group_name"`
	Item      interface{} `json:"item"`
	Priority  int64       `json:"priority"`
}

type TodoUpdateRequest struct {
	ID      int64  `json:"id" validate:"required,number,gte=1"`
	Name    string `json:"name" validate:"required"`
	UserID  int    `json:"user_id" validate:"required"`
	GroupID int    `json:"group_id" validate:"required"`
}

type TodoDeleteRequest struct {
	ID int64 `json:"id" validate:"required,number,gte=1"`
}
