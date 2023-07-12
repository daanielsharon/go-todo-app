package service

import (
	"context"
	"server/model/web"
)

type TodoService interface {
	CreateTodo(c context.Context, req *web.TodoCreateRequest) *web.TodoCreateResponse
	GetTodoByUsername(c context.Context, req *web.TodoGetRequest) *[]web.TodoGetResponse
	RemoveTodo(c context.Context, req *web.TodoDeleteRequest)
}
