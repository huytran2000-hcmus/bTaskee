package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/pkg/validator"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/repository"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrFailedValidation  = errors.New("failed validation")
	ErrFailedUpstreamReq = errors.New("failed to call upstream service")
)

type Task interface {
	CreateOne(ctx context.Context, task *model.CreateTaskRequest) (uuid.UUID, error)
	GetOne(ctx context.Context, id uuid.UUID) (*model.Task, error)
}

type task struct {
	repo        repository.Repository
	pricingRepo repository.PricingRepository
	sendRepo    repository.SendRepository
	validator   *validator.Validator
}

func NewTask(repo repository.Repository, pricingRepo repository.PricingRepository, sendRepo repository.SendRepository) *task {
	return &task{
		repo:        repo,
		pricingRepo: pricingRepo,
		sendRepo:    sendRepo,
		validator:   validator.New(),
	}
}

func (s task) CreateOne(ctx context.Context, req *model.CreateTaskRequest) (uuid.UUID, error) {
	err := s.validator.Validate(req)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%w: %w", ErrFailedValidation, err)
	}
	task := model.CreateTaskRequestToTask(req)
	task.ID = uuid.New()

	calculatePriceReq := model.CalculatePriceRequest{
		HouseType:    task.WorkingDetails.HouseType,
		ServiceTypes: task.WorkingDetails.ServiceTypes,
		FromTime:     task.WorkingDetails.FromTime,
		ToTime:       task.WorkingDetails.ToTime,
	}

	calculatePriceRes, err := s.pricingRepo.GetPricing(ctx, calculatePriceReq)
	if err != nil {
		return uuid.Nil, fmt.Errorf("pricing: %w: %w", ErrFailedUpstreamReq, err)
	}
	task.Total = calculatePriceRes.Total
	fmt.Printf("total: %+v\n", calculatePriceRes.Total)

	err = s.repo.InsertTask(ctx, task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert task into repository: %w", err)
	}

	sendTask := model.ToSendTaskRequest(task)
	_, err = s.sendRepo.SendTask(ctx, sendTask)
	if err != nil {
		return uuid.Nil, fmt.Errorf("send: %w: %w", ErrFailedUpstreamReq, err)
	}

	return task.ID, err
}

func (s task) GetOne(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}
