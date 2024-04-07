package restful

import (
	"net/http"

	"github.com/huytran2000-hcmus/bTaskee/send/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/service"
	echo "github.com/labstack/echo/v4"
)

type sendHandler struct {
	svc service.Send
}

func NewSendHandler(svc service.Send) *sendHandler {
	return &sendHandler{
		svc: svc,
	}
}

func (h *sendHandler) send(c echo.Context) error {
	var req model.SendTaskRequest

	err := c.Bind(&req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.svc.Send(c.Request().Context(), req)
	if err != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, err)
	}

	return sendSuccessResponse(c, http.StatusOK, res, "")
}
