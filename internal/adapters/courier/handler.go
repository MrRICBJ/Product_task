package courier

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sss/internal/adapters"
	"sss/internal/apperror"
	"sss/internal/courier"
	"sss/pkg/lib"
	"strconv"
)

const (
	couriersURL  = "/couriers"
	courierIdURL = "/couriers/:courier_id"
)

type handler struct {
	repo courier.Repository
}

func NewHandler(repo courier.Repository) adapters.Handler {
	return &handler{repo: repo}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(couriersURL, h.GetAll)
	router.GET(courierIdURL, h.GetById)
	router.POST(couriersURL, h.Create)
}

func (h *handler) GetAll(c echo.Context) error {
	limit, offset := lib.GetLimOff(c)
	statusCode, o := h.repo.GetAll(context.Background(), limit, offset)
	return c.JSON(statusCode, o)
}

func (h *handler) GetById(c echo.Context) error {
	tmpId, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetById(context.Background(), tmpId)
	return c.JSON(statusCode, o)
}

func (h *handler) Create(c echo.Context) error {
	courierReq := courier.CreateCourierRequest{}
	err := c.Bind(&courierReq)
	if err != nil {
		log.Fatal(err)
	}

	statusCode, o := h.repo.Create(context.Background(), &courierReq)
	return c.JSON(statusCode, o)
}
