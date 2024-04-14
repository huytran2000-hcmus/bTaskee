package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/huytran2000-hcmus/bTaskee/pricing/internal/model"
)

const (
	TwoBedRoomPricePerHour  = 100
	ThreeBedRoomPriceByHour = 150
	FourBedRoomPriceByHour  = 200

	WeekdayRate = 1
	WeekendRate = 1.5
)

var (
	ErrWorkingsNotInTheSameDate = errors.New("working time span is not in a day")
	ErrWorkingLessThanMinimum   = errors.New("working hours less than minimum hours")
	ErrWorkingMoreThanMaximum   = errors.New("working hours more than maximum hours")
	ErrHouseTypeNotFound        = errors.New("house type not found")
	ErrFailedValidation         = errors.New("failed validation")
)

var (
	MinimumWorkingHours = 2 * time.Hour
	MaximumWorkingHours = 4 * time.Hour
)

type Pricing interface {
	GetPrice(ctx context.Context, req model.CalculatePriceRequest) (float64, error)
}

func NewPricing() *pricing {
	return &pricing{
		validator: validator.New(),
	}
}

type pricing struct {
	validator *validator.Validate
}

func (s pricing) GetPrice(ctx context.Context, req model.CalculatePriceRequest) (float64, error) {
	var res float64
	err := s.validator.Struct(req)
	if err != nil {
		return res, fmt.Errorf("%w: %w", ErrFailedValidation, err)
	}

	toDate := req.ToTime.Truncate(24 * time.Hour)
	fromDate := req.FromTime.Truncate(24 * time.Hour)

	if !toDate.Equal(fromDate) {
		return res, ErrWorkingsNotInTheSameDate
	}

	workingHours := req.ToTime.Sub(req.FromTime)
	if workingHours < MinimumWorkingHours {
		return res, ErrWorkingLessThanMinimum
	}

	if workingHours > MaximumWorkingHours {
		return res, ErrWorkingMoreThanMaximum
	}

	var rate float64
	weekday := toDate.Weekday()
	switch weekday {
	case time.Saturday, time.Sunday:
		rate = WeekendRate
	default:
		rate = WeekdayRate
	}

	var pricePerHour float64
	switch req.HouseType {
	case model.TwoRoom:
		pricePerHour = TwoBedRoomPricePerHour
	case model.ThreeRoom:
		pricePerHour = ThreeBedRoomPriceByHour
	case model.FourRoom:
		pricePerHour = FourBedRoomPriceByHour
	default:
		return res, ErrHouseTypeNotFound
	}

	res = rate * pricePerHour * float64(workingHours/time.Hour)

	return res, nil
}
