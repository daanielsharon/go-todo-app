package web

type UserGetUsernameRequest struct {
	Username string `json:"username" validate:"required,min=1,max=30"`
}

type UserGetUsernameResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserCreateUsernameRequest struct {
	Username string `json:"username" validate:"required,min=1,max=30"`
}

type UserCreateUsernameResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
