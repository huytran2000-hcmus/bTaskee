package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/repository"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Task interface {
	CreateOne(ctx context.Context, task *model.CreateTaskRequest) (id uuid.UUID, err error)
	GetOne(ctx context.Context, id uuid.UUID) (*model.Task, error)
}

type task struct {
	repo repository.Repository
}

func NewTask(repo repository.Repository) *task {
	return &task{
		repo: repo,
	}
}

func (s task) CreateOne(ctx context.Context, req *model.CreateTaskRequest) (id uuid.UUID, err error) {
	task := model.CreateTaskRequestToTask(req)
	task.ID = uuid.New()
	err = s.repo.InsertTask(ctx, task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert task into repository: %w", err)
	}

	return task.ID, err
}

func (s task) GetOne(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}
