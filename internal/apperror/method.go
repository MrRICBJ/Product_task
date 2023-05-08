package apperror

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

const (
	limit  = "limit"
	offset = "offset"
)

func GetLimOff(c echo.Context) (int32, int32) {
	limit, err := strconv.Atoi(c.QueryParam(limit))
	if err != nil {
		limit = 1
	}

	offset, err := strconv.Atoi(c.QueryParam(offset))
	if err != nil {
		offset = 0
	}

	return int32(limit), int32(offset)
}
