package web

type TodoCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	UserID  int    `json:"userId" validate:"required"`
	GroupID int    `json:"groupId" validate:"required"`
}

type TodoCreateUpdateResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	UserID  int    `json:"userId"`
	GroupID int    `json:"groupId"`
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
	UserID  int    `json:"userId" validate:"required"`
	GroupID int    `json:"groupId" validate:"required"`
}

type TodoUpdatePriority struct {
	OriginID       int64 `json:"originId" validate:"required,number,gte=1"`
	OriginPriority int64 `json:"originPriority" validate:"required,number,gte=1"`
	TargetID       int64 `json:"targetId" validate:"required,number,gte=1"`
	TargetPriority int64 `json:"targetPriority" validate:"required,number,gte=1"`
}

type TodoDeleteRequest struct {
	ID int64 `json:"id" validate:"required,number,gte=1"`
}
