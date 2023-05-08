package courier

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sss/internal/adapters"
	"sss/internal/apperror"
	"sss/internal/courier"
	"strconv"
	"time"
)

const (
	couriersURL  = "/couriers"
	courierIdURL = "/couriers/:courier_id"
	courierMeta  = "/couriers/meta-info/:courier_id"
	start        = "startDate"
	//start        = "start_date"
	end = "endDate"
	//end          = "end_date"
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
	router.GET(courierMeta, h.GetMetaInf)
	router.POST(couriersURL, h.Create)
}

func (h *handler) GetAll(c echo.Context) error {
	limit, offset := apperror.GetLimOff(c)
	statusCode, o := h.repo.GetAll(context.Background(), limit, offset)
	return c.JSON(statusCode, o)
}

func (h *handler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetById(context.Background(), id)
	return c.JSON(statusCode, o)
}

func (h *handler) GetMetaInf(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("courier_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	start, end, err := getStartEnd(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetMetaInf(context.Background(), id, *start, *end)
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

func getStartEnd(c echo.Context) (*time.Time, *time.Time, error) {
	startStr := c.QueryParam(start)
	endStr := c.QueryParam(end)

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return nil, nil, err
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		return nil, nil, err
	}

	end = end.Add(time.Hour * 23)
	end = end.Add(time.Minute * 59)
	end = end.Add(time.Second * 59)
	return &start, &end, nil
}
