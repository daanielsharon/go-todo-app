package itemserv

import (
	"context"
	"server/model/web"
)

type ItemService interface {
	Create(c context.Context, req *web.TodoCreateRequest) *web.TodoCreateUpdateResponse
	Update(c context.Context, req *web.TodoUpdateRequest) *web.TodoCreateUpdateResponse
	GetByUsername(c context.Context, req *web.TodoGetRequest) *[]web.TodoGetResponse
	Remove(c context.Context, req *web.TodoDeleteRequest)
}
