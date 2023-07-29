package web

type UserGetRequest struct {
	Username string `json:"username" validate:"required,min=1,max=30"`
	Password string `json:"password" validate:"required,min=1"`
}

type UserGetResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=30"`
	Password string `json:"password" validate:"required,min=1"`
}

type UserCreateResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
