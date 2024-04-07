package service

import (
	"context"

	"github.com/huytran2000-hcmus/bTaskee/send/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/repository"
)

type Send interface {
	Send(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error)
}

func NewSend(emailRepo repository.Emai) *send {
	return &send{
		emailRepo: emailRepo,
	}
}

type send struct {
	emailRepo repository.Emai
}

func (s send) Send(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error) {
	ids, err := s.emailRepo.SendEmail(ctx, req)

	return model.SendTaskResponse{
		EmailIDs: ids,
	}, err
}
