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

// SendTask send task to tasker
// @Summary Send task to tasker
// @Description Send task to tasker
// @Tags Task
// @Accept  json
// @Produce  json
// @Param   request body model.SendTaskRequest true "Send Task Request"
// @Success 200 {object} successResponse{data=model.SendTaskResponse} "success"
// @Failure 400 {object} errorResponse "bad request"
// @Router /api/v1/task:send [post]
func (h *sendHandler) sendTask(c echo.Context) error {
	var req model.SendTaskRequest

	err := c.Bind(&req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.svc.SendTask(c.Request().Context(), req)
	if err != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, err)
	}

	return sendSuccessResponse(c, http.StatusOK, res, "")
}
