package service

import (
	"context"
	"server/model/web"
)

type UserService interface {
	CreateUsername(c context.Context, req *web.UserCreateUsernameRequest) *web.UserCreateUsernameResponse
	GetUsername(c context.Context, req *web.UserGetUsernameRequest) *web.UserGetUsernameResponse
}
