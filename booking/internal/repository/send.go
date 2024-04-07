package repository

import (
	"context"

	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
)

type SendRepository interface {
	SendTask(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error)
}
