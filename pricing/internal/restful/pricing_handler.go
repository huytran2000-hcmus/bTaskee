package restful

import (
	"net/http"

	"github.com/huytran2000-hcmus/bTaskee/pricing/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/pricing/internal/service"
	"github.com/labstack/echo/v4"
)

type pricingHandler struct {
	svc service.Pricing
}

func NewPricingHandler(svc service.Pricing) *pricingHandler {
	return &pricingHandler{
		svc: svc,
	}
}

func (h *pricingHandler) calculatePricing(c echo.Context) error {
	var req model.CalculatePriceRequest

	err := c.Bind(&req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.svc.GetPricing(c.Request().Context(), req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	return sendSuccessResponse(c, http.StatusOK, res, "")
}
