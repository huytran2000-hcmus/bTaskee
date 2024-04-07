package service

import (
	"context"
	"errors"
	"time"

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
)

var (
	MinimumWorkingHours = 2 * time.Hour
	MaximumWorkingHours = 4 * time.Hour
)

type Pricing interface {
	GetPricing(ctx context.Context, req model.CalculatePriceRequest) (model.PricingResponse, error)
}

func NewPricing() *pricing {
	return &pricing{}
}

type pricing struct{}

func (s pricing) GetPricing(ctx context.Context, req model.CalculatePriceRequest) (model.PricingResponse, error) {
	var res model.PricingResponse
	toDate := req.To.Truncate(24 * time.Hour)
	fromDate := req.From.Truncate(24 * time.Hour)

	if !toDate.Equal(fromDate) {
		return res, ErrWorkingsNotInTheSameDate
	}

	workingHours := req.From.Sub(req.To)
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

	res.Total = rate * pricePerHour * float64(workingHours/time.Hour)

	return res, nil
}
