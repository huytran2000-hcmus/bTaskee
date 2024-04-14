package restful

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func parseUUIDFromURI(c echo.Context) (uuid.UUID, error) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return uuid.Nil, err
	}

	if id == uuid.Nil {
		return uuid.Nil, errors.New("empty id")
	}

	return id, nil
}
