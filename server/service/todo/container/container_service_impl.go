package containerserv

import (
	"context"
	"database/sql"
	"server/exception"
	"server/model/domain"
	"server/model/web"
	containerrepo "server/repository/todo/container"
	"time"

	"github.com/go-playground/validator/v10"
)

type ContainerServiceImpl struct {
	ContainerRepository containerrepo.ContainerRepository
	DB                  *sql.DB
	Validate            *validator.Validate
	Timeout             time.Duration
}

func NewContainerService(containerRepository containerrepo.ContainerRepository, db *sql.DB, validator *validator.Validate, timeout time.Duration) ContainerService {
	return &ContainerServiceImpl{
		ContainerRepository: containerRepository,
		DB:                  db,
		Validate:            validator,
		Timeout:             timeout,
	}
}

func (s *ContainerServiceImpl) Create(c context.Context, req *web.ContainerCreateRequest) *web.ContainerCreateResponse {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	// find the total container, if it exceeds 5, reject
	container := domain.Container{
		UserId: req.UserId,
	}

	total := s.ContainerRepository.FindTotalContainer(ctx, s.DB, &container)
	if *total >= 4 {
		panic(exception.NewValidationError("cannot create more than 4 containers!"))
	}

	containerReq := &domain.Container{
		UserId:    req.UserId,
		GroupName: req.GroupName,
		Priority:  req.Priority,
	}

	result := s.ContainerRepository.Save(ctx, s.DB, containerReq)

	res := web.ContainerCreateResponse{
		ID:        result.ID,
		UserId:    result.UserId,
		GroupName: result.GroupName,
		Priority:  result.Priority,
	}

	return &res
}

func (s *ContainerServiceImpl) UpdatePriority(c context.Context, req *web.TodoUpdatePriority) *web.TodoUpdatePriority {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	// swap priority between two

	origin := &domain.TodoPriority{
		ID:       req.OriginID,
		Priority: int(req.OriginPriority),
	}

	resOrigin := s.ContainerRepository.UpdatePriority(ctx, s.DB, origin)

	target := &domain.TodoPriority{
		ID:       req.TargetID,
		Priority: int(req.TargetPriority),
	}

	resTarget := s.ContainerRepository.UpdatePriority(ctx, s.DB, target)

	res := &web.TodoUpdatePriority{
		OriginID:       resOrigin.ID,
		OriginPriority: int64(resOrigin.Priority),
		TargetID:       resTarget.ID,
		TargetPriority: int64(resTarget.Priority),
	}

	return res
}
