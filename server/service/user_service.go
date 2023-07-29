package service

import (
	"context"
	"server/model/web"
)

type UserService interface {
	Create(c context.Context, req *web.UserCreateRequest) *web.UserCreateResponse
	Get(c context.Context, req *web.UserGetRequest) *web.UserGetResponse
}
