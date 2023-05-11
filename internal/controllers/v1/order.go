package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"sss/internal/apperror"
	"sss/internal/controllers/dto"
	"sss/internal/controllers/v1/utils"
	"sss/internal/service"
	"strconv"
)

const (
	ordersURL    = "/orders"
	orderIdURL   = "/orders/:order_id"
	ordersComURL = "/orders/complete"
	limit        = "limit"
	offset       = "offset"
	orderId      = "order_id"
)

type handler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(ordersURL, h.getOrders)
	router.GET(orderIdURL, h.getOrder)
	router.POST(ordersURL, h.createOrder)
	router.POST(ordersComURL, h.completeOrder)
}

func (h *handler) getOrders(c echo.Context) error {
	limit, offset, err := utils.GetLimOff(c.QueryParam(limit), c.QueryParam(offset))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	result, err := h.service.GetOrders(context.Background(), limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}
	return c.JSON(http.StatusOK, result)
}

func (h *handler) getOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param(orderId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	result, err := h.service.GetOrder(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, apperror.NotFoundResponse{})
	}
	return c.JSON(http.StatusOK, result)
}

func (h *handler) createOrder(c echo.Context) error {
	orderReq := dto.CreateOrderRequest{}

	err := c.Bind(&orderReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	if err := utils.Validation(&orderReq); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	result, err := h.service.CreateOrders(context.Background(), &orderReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}
	return c.JSON(http.StatusOK, result)
}

func (h *handler) completeOrder(c echo.Context) error {
	orderReq := dto.CompleteOrderRequestDto{}

	err := c.Bind(&orderReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	result, err := h.service.CompleteOrders(context.Background(), orderReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}
	return c.JSON(http.StatusOK, result)
}
