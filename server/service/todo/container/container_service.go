package containerserv

import (
	"context"
	"server/model/web"
)

type ContainerService interface {
	UpdatePriority(c context.Context, req *web.TodoUpdatePriority) *web.TodoUpdatePriority
}
