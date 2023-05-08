package order

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sss/internal/adapters"
	"sss/internal/apperror"
	"sss/internal/order"
	"strconv"
)

const (
	ordersURL    = "/orders"
	orderIdURL   = "/orders/:order_id"
	ordersComURL = "/orders/complete"
)

type handler struct {
	repo order.Repository
}

func NewHandler(repo order.Repository) adapters.Handler {
	return &handler{repo: repo}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(ordersURL, h.GetAll)
	router.GET(orderIdURL, h.GetById)
	router.POST(ordersURL, h.Create)
	router.POST(ordersComURL, h.UpdateComplete)
}

func (h *handler) GetAll(c echo.Context) error {
	limit, offset := apperror.GetLimOff(c)
	statusCode, o := h.repo.GetAll(context.Background(), limit, offset)
	return c.JSON(statusCode, o)
}

func (h *handler) GetById(c echo.Context) error {
	tmpId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetById(context.Background(), tmpId)
	return c.JSON(statusCode, o)
}

func (h *handler) Create(c echo.Context) error {
	orderReq := order.CreateOrderRequest{}

	err := c.Bind(&orderReq)
	if err != nil {
		log.Fatal(err)
	}

	statusCode, o := h.repo.Create(context.Background(), &orderReq)
	return c.JSON(statusCode, o)
}

func (h *handler) UpdateComplete(c echo.Context) error {
	orderReq := order.CompleteOrderRequestDto{}

	err := c.Bind(&orderReq)
	if err != nil {
		log.Fatal(err)
	}

	statusCode, o := h.repo.Update(context.Background(), &orderReq)
	return c.JSON(statusCode, o)
}
