package restful

import "github.com/labstack/echo/v4"

type successResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func sendSuccessResponse(c echo.Context, code int, data any, mess string) error {
	return c.JSON(code, successResponse{
		Data:    data,
		Message: mess,
	})
}

func sendErrorResponse(c echo.Context, code int, err error) error {
	return c.JSON(code, errorResponse{
		Error: err.Error(),
	})
}
