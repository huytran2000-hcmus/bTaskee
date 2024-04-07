package restful

import (
	"errors"
	"net/http"

	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/service"
	"github.com/labstack/echo/v4"
)

type taskHandler struct {
	svc service.Task
}

func newTaskHandler(svc service.Task) *taskHandler {
	return &taskHandler{
		svc: svc,
	}
}

func (h *taskHandler) createTaskHandler(c echo.Context) error {
	req := &model.CreateTaskRequest{}

	err := c.Bind(req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	id, err := h.svc.CreateOne(c.Request().Context(), req)
	if err != nil {
		return sendErrorResponse(c, http.StatusInternalServerError, err)
	}

	return sendSuccessResponse(c, http.StatusOK, id, "successly create task")
}

func (h *taskHandler) getTaskByID(c echo.Context) error {
	id, err := parseIDFromURI(c)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	task, err := h.svc.GetOne(c.Request().Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return sendErrorResponse(c, http.StatusNotFound, err)
		default:
			return sendErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return sendSuccessResponse(c, http.StatusOK, task, "")
}
