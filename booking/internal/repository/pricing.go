package repository

import (
	"context"

	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
)

type PricingRepository interface {
	GetPricing(ctx context.Context, req model.CalculatePriceRequest) (model.CalculatePriceResponse, error)
}
