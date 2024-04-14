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

// CalculatePrice calculate price of a task
// @Summary Calculate price of a task
// @Description Calculate price of a task
// @Tags Task
// @Accept  json
// @Produce  json
// @Param   request body model.CalculatePriceRequest true "Calculate Price Request"
// @Success 200 {object} successResponse{data=float32} "success"
// @Failure 400 {object} errorResponse "bad request"
// @Router /api/v1/task:calculate-price [post]
func (h *pricingHandler) calculatePrice(c echo.Context) error {
	var req model.CalculatePriceRequest

	err := c.Bind(&req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	price, err := h.svc.GetPrice(c.Request().Context(), req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	res := model.CalculatePriceResponse{
		Total: price,
	}
	return sendSuccessResponse(c, http.StatusOK, res, "")
}
