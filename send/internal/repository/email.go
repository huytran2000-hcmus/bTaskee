package repository

import (
	"context"

	"github.com/huytran2000-hcmus/bTaskee/send/internal/model"
)

type Emai interface {
	SendEmail(ctx context.Context, req model.SendTaskRequest) (ids []string, err error)
	// GetEmails(ctx context.Context, ids ...string)
}
