package web

type ContainerCreateRequest struct {
	UserId    int64  `json:"userId" validate:"required,number,gte=1"`
	GroupName string `json:"groupName" validate:"required"`
	Priority  uint8  `json:"priority" validate:"required,number,gte=1"`
}

type ContainerCreateResponse struct {
	ID        int64  `json:"id" validate:"required,number,gte=1"`
	UserId    int64  `json:"userId" validate:"required,number,gte=1"`
	GroupName string `json:"groupName" validate:"required"`
	Priority  uint8  `json:"priority" validate:"required,number,gte=1"`
}

type ContainerDeleteRequest struct {
	ID int64 `json:"id" validate:"required,number,gte=1"`
}
