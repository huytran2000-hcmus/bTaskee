package pricing_api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/huytran2000-hcmus/bTaskee/booking/config"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/pkg/httpapi"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/repository"
)

var _ repository.PricingRepository = (*PricingAPI)(nil)

type PricingAPI struct{}

func New() *PricingAPI {
	return &PricingAPI{}
}

func (p *PricingAPI) GetPricing(ctx context.Context, req model.CalculatePriceRequest) (model.CalculatePriceResponse, error) {
	httpReq := httpapi.HTTPRequest{
		Method: http.MethodPost,
		URL:    config.LoadEnv().CalcalatePriceURI,
		Body:   req,
	}

	res, err := httpapi.SendHTTPRequest[model.CalculatePricePayload](ctx, httpReq, httpapi.HTTPOptions{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return res.Data, err
	}

	if res.Error != "" {
		return res.Data, errors.New(res.Error)
	}

	return res.Data, err
}
