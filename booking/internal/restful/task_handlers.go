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

// CreateTask creates a new task.
// @Summary Create a new task
// @Description Create a new task
// @Tags Task
// @Accept  json
// @Produce  json
// @Param   request body model.CreateTaskRequest true "Create Task Request"
// @Success 200 {object} successResponse{data=string} "success"
// @Failure 400 {object} errorResponse{data=string} "bad request"
// @Router /api/v1/tasks [post]
func (h *taskHandler) createTaskHandler(c echo.Context) error {
	req := &model.CreateTaskRequest{}

	err := c.Bind(req)
	if err != nil {
		return sendErrorResponse(c, http.StatusBadRequest, err)
	}

	id, err := h.svc.CreateOne(c.Request().Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrFailedValidation):
			return sendErrorResponse(c, http.StatusBadRequest, err)
		default:
			return sendErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return sendSuccessResponse(c, http.StatusOK, id, "successly create task")
}

// GetTaskByID get a task by id
// @Summary Get a task by id
// @Description Get a task by id
// @Tags Task
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} successResponse{data=model.Task} "success"
// @Failure 404 {object} errorResponse{data=string} "bad request"
// @Router /api/v1/tasks/{id} [get]
func (h *taskHandler) getTaskByID(c echo.Context) error {
	id, err := parseUUIDFromURI(c)
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
