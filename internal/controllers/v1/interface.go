package v1

import "github.com/labstack/echo/v4"

type Handler interface {
	Register(router *echo.Echo)
}
