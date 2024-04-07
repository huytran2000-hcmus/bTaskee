package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/repository"
)

var (
	ErrFailedValidation = errors.New("failed validation")
)

type Send interface {
	SendTask(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error)
}

func NewSend(emailRepo repository.Emai) *send {
	return &send{
		emailRepo: emailRepo,
		validator: validator.New(),
	}
}

type send struct {
	emailRepo repository.Emai
	validator *validator.Validate
}

func (s send) SendTask(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error) {
	var res model.SendTaskResponse
	err := s.validator.Struct(req)
	if err != nil {
		return res, fmt.Errorf("%w: %w", ErrFailedValidation, err)
	}
	ids, err := s.emailRepo.SendEmail(ctx, req)
	res.MessageIDs = ids

	return res, err
}
