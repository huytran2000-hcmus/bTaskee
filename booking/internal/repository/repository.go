package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
)

type Repository interface {
	InsertTask(ctx context.Context, task *model.Task) error
	GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
}
