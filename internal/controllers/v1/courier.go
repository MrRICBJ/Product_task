package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sss/internal/apperror"
	"sss/internal/entity"
	courier2 "sss/internal/repository/courier"
	"strconv"
	"time"
)

const (
	couriersURL  = "/couriers"
	courierIdURL = "/couriers/:courier_id"
	courierMeta  = "/couriers/meta-info/:courier_id"
	start        = "startDate"
	courId       = "courier_id"
	//start        = "start_date"
	end = "endDate"
	//end          = "end_date"
	limit  = "limit"
	offset = "offset"
)

type handler struct {
	repo courier2.CourRepo
}

func NewHandler(repo courier2.CourRepo) handler.Handler {
	return &handler{repo: repo}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(couriersURL, h.GetAll)
	router.GET(courierIdURL, h.GetById)
	router.GET(courierMeta, h.GetMetaInf)
	router.POST(couriersURL, h.Create)
}

func (h *handler) GetAll(c echo.Context) error {
	limit, offset, err := handler.GetLimOff(c.QueryParam(limit), c.QueryParam(offset))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetAll(context.Background(), limit, offset)
	return c.JSON(statusCode, o)
}

func (h *handler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param(courId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetById(context.Background(), id)
	return c.JSON(statusCode, o)
}

func (h *handler) GetMetaInf(c echo.Context) error {
	id, err := strconv.Atoi(c.Param(courId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	start, end, err := getStartEnd(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	statusCode, o := h.repo.GetMetaInf(context.Background(), id, start, end)
	return c.JSON(statusCode, o)
}

func (h *handler) Create(c echo.Context) error {
	courierReq := entity.CreateCourierRequest{}
	err := c.Bind(&courierReq)
	if err != nil {
		log.Fatal(err)
	}

	statusCode, o := h.repo.Create(context.Background(), &courierReq)
	return c.JSON(statusCode, o)
}

func getStartEnd(c echo.Context) (time.Time, time.Time, error) {
	startStr := c.QueryParam(start)
	endStr := c.QueryParam(end)

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	//end = end.Add(time.Hour * 23)
	//end = end.Add(time.Minute * 59)
	//end = end.Add(time.Second * 59)
	return start, end, nil
}