package itemserv

import (
	"context"
	"database/sql"
	"server/exception"
	"server/helper"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	containerrepo "server/repository/todo/container"
	itemrepo "server/repository/todo/item"
	"time"

	"github.com/go-playground/validator/v10"
)

type ItemServiceImpl struct {
	ItemRepository      itemrepo.ItemRepository
	ContainerRepository containerrepo.ContainerRepository
	UserRepository      repository.UserRepository
	DB                  *sql.DB
	Timeout             time.Duration
	Validate            *validator.Validate
}

func NewItemService(itemRepository itemrepo.ItemRepository, containerRepository containerrepo.ContainerRepository, userRepository repository.UserRepository, db *sql.DB, validator *validator.Validate, timeout time.Duration) ItemService {
	return &ItemServiceImpl{
		ItemRepository:      itemRepository,
		ContainerRepository: containerRepository,
		UserRepository:      userRepository,
		DB:                  db,
		Timeout:             timeout,
		Validate:            validator,
	}
}

func (s *ItemServiceImpl) Create(c context.Context, req *web.TodoCreateRequest) *web.TodoCreateUpdateResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	newUser := &domain.TodoListInsertUpdate{
		Name:    req.Name,
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}

	r, err := s.ItemRepository.Save(ctx, s.DB, newUser)
	if err != nil {
		panic(err)
	}

	res := &web.TodoCreateUpdateResponse{
		ID:      r.ID,
		Name:    r.Name,
		UserID:  r.UserID,
		GroupID: r.GroupID,
	}

	return res
}

func (s *ItemServiceImpl) GetByUsername(c context.Context, req *web.TodoGetRequest) *[]web.TodoGetResponse {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	_, err := s.UserRepository.Find(ctx, s.DB, &user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	res, err := s.ItemRepository.FindByUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	var response []web.TodoGetResponse

	for _, val := range *res {
		var item interface{}

		// if first data is nil, it means there's no data for the entire todo group
		if val.Item[0] == nil {
			item = make([]interface{}, 0)
		} else {
			item = val.Item
		}

		response = append(response, web.TodoGetResponse{
			ID:        val.ID,
			GroupName: val.Name,
			Item:      item,
			Priority:  val.Priority,
		})
	}

	return &response
}

func (s *ItemServiceImpl) Update(c context.Context, req *web.TodoUpdateRequest) *web.TodoCreateUpdateResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	_, err = s.ItemRepository.FindById(ctx, s.DB, req.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = s.ItemRepository.FindByName(ctx, s.DB, req.Name)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	groupReq := &domain.TodoGroup{
		ID:     int64(req.GroupID),
		UserID: req.UserID,
	}

	_, err = s.ContainerRepository.FindGroup(ctx, s.DB, groupReq)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	data := &domain.TodoListInsertUpdate{
		ID:      req.ID,
		Name:    req.Name,
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}

	r := s.ItemRepository.Update(ctx, s.DB, data)

	res := &web.TodoCreateUpdateResponse{
		ID:      r.ID,
		Name:    r.Name,
		GroupID: r.GroupID,
		UserID:  r.UserID,
	}

	return res
}

func (s *ItemServiceImpl) Remove(c context.Context, req *web.TodoDeleteRequest) {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	_, err = s.ItemRepository.FindById(ctx, s.DB, req.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.ItemRepository.Delete(ctx, s.DB, &domain.TodoList{ID: req.ID})
	if err != nil {
		panic(err)
	}
}
