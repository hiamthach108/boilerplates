package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/hiamthach108/dreon-backend-service/internal/aggregate"
	"github.com/hiamthach108/dreon-backend-service/internal/errorx"
	"github.com/hiamthach108/dreon-backend-service/internal/model"
	"github.com/hiamthach108/dreon-backend-service/internal/repository"
	"github.com/hiamthach108/dreon-backend-service/pkg/validator"
)

type IExampleSvc interface {
	GetAll(ctx context.Context) ([]aggregate.ExampleAggregate, error)
	Create(ctx context.Context, req *aggregate.CreateExampleReq) (*aggregate.ExampleAggregate, error)
	Update(ctx context.Context, id string, req *aggregate.UpdateExampleReq) error
}

type ExampleSvc struct {
	repo repository.IExampleRepository
}

func NewExampleSvc(repo repository.IExampleRepository) IExampleSvc {
	return &ExampleSvc{repo: repo}
}

func (s *ExampleSvc) GetAll(ctx context.Context) ([]aggregate.ExampleAggregate, error) {
	rows, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, errorx.Wrap(errorx.ErrInternal, fmt.Errorf("list examples: %w", err))
	}
	out := make([]aggregate.ExampleAggregate, 0, len(rows))
	for i := range rows {
		var agg aggregate.ExampleAggregate
		agg.FromModel(&rows[i])
		out = append(out, agg)
	}
	return out, nil
}

func (s *ExampleSvc) Create(ctx context.Context, req *aggregate.CreateExampleReq) (*aggregate.ExampleAggregate, error) {
	if err := validator.ValidateStruct(req); err != nil {
		return nil, errorx.Wrap(errorx.ErrBadRequest, validator.FormatValidationError(err))
	}
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errorx.New(errorx.ErrBadRequest, "name is required")
	}
	if existing := s.repo.FindByName(ctx, name); existing != nil {
		return nil, errorx.New(errorx.ErrConflict, "example name already exists")
	}

	example := &model.Example{
		Name:        name,
		Description: req.Description,
		IsActive:    true,
	}
	if req.IsActive != nil {
		example.IsActive = *req.IsActive
	}

	created, err := s.repo.Create(ctx, example)
	if err != nil {
		return nil, errorx.Wrap(errorx.ErrInternal, fmt.Errorf("create example: %w", err))
	}
	if created == nil || created.ID == "" {
		return nil, errorx.New(errorx.ErrInternal, "example created but ID is empty")
	}
	var agg aggregate.ExampleAggregate
	agg.FromModel(created)
	return &agg, nil
}

func (s *ExampleSvc) Update(ctx context.Context, id string, req *aggregate.UpdateExampleReq) error {
	if strings.TrimSpace(id) == "" {
		return errorx.New(errorx.ErrBadRequest, "id is required")
	}
	if err := validator.ValidateStruct(req); err != nil {
		return errorx.Wrap(errorx.ErrBadRequest, validator.FormatValidationError(err))
	}
	current := s.repo.FindOneById(ctx, id)
	if current == nil {
		return errorx.New(errorx.ErrNotFound, "example not found")
	}

	updates := model.Example{}
	var fields []string

	if req.Name != nil {
		newName := strings.TrimSpace(*req.Name)
		if newName == "" {
			return errorx.New(errorx.ErrBadRequest, "name cannot be empty")
		}
		if other := s.repo.FindByName(ctx, newName); other != nil && other.ID != id {
			return errorx.New(errorx.ErrConflict, "example name already exists")
		}
		updates.Name = newName
		fields = append(fields, "Name")
	}
	if req.Description != nil {
		updates.Description = *req.Description
		fields = append(fields, "Description")
	}
	if req.IsActive != nil {
		updates.IsActive = *req.IsActive
		fields = append(fields, "IsActive")
	}
	if len(fields) == 0 {
		return errorx.New(errorx.ErrBadRequest, "no fields to update")
	}

	if err := s.repo.Update(ctx, id, updates, fields...); err != nil {
		return errorx.Wrap(errorx.ErrInternal, fmt.Errorf("update example: %w", err))
	}
	return nil
}
