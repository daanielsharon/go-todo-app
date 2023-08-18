package containerserv

import (
	"context"
	"server/model/web"
)

type ContainerService interface {
	Create(c context.Context, req *web.ContainerCreateRequest) *web.ContainerCreateResponse
	UpdatePriority(c context.Context, req *web.TodoUpdatePriority) *web.TodoUpdatePriority
	Delete(c context.Context, req *web.ContainerDeleteRequest)
}
