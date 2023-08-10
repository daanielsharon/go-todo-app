package containerserv

import (
	"context"
	"server/model/web"
)

type ContainerService interface {
	CreateContainer(c context.Context, req *web.ContainerCreateRequest) *web.ContainerCreateResponse
	UpdatePriority(c context.Context, req *web.TodoUpdatePriority) *web.TodoUpdatePriority
}
