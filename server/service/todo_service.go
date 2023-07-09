package service

import (
	"context"
	"server/model/web"
)

type TodoService interface {
	CreateTodo(c context.Context, req *web.TodoCreateRequest) (*web.TodoCreateResponse, error)
	RemoveTodo(c context.Context, req *web.TodoDeleteRequest) error
}
