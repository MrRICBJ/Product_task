package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sss/internal/apperror"
	"sss/internal/controllers/dto"
	"sss/internal/controllers/v1/utils"
	"sss/internal/service"
	"strconv"
	"time"
)

const (
	couriersURL  = "/couriers"
	courierIdURL = "/couriers/:courier_id"
	courierMeta  = "/couriers/meta-info/:courier_id"
	start        = "startDate"
	courId       = "courier_id"
	end          = "endDate"
)

type courHandler struct {
	service service.CourService
}

func NewCourHandler(service service.CourService) Handler {
	return &courHandler{service: service}
}

func (h *courHandler) Register(router *echo.Echo) {
	router.GET(couriersURL, h.getCouriers)
	router.GET(courierIdURL, h.getCourierById)
	router.GET(courierMeta, h.getCourierMetaInfo)
	router.POST(couriersURL, h.createCourier)
}

func (h *courHandler) getCouriers(c echo.Context) error {
	limit, offset, err := utils.GetLimOff(c.QueryParam(limit), c.QueryParam(offset))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	result, err := h.service.GetCouriers(context.Background(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *courHandler) getCourierById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param(courId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	result, err := h.service.GetCourierById(context.Background(), int64(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, apperror.NotFoundResponse{})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *courHandler) createCourier(c echo.Context) error {
	courierReq := dto.CreateCourierRequest{}

	err := c.Bind(&courierReq)
	if err != nil {
		log.Fatal(err)
	}

	result, err := h.service.CreateCourier(context.Background(), &courierReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *courHandler) getCourierMetaInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param(courId))
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	//}

	start, end, _ := getStartEnd(c)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, apperror.BadRequestResponse{})
	//}

	result, _ := h.service.GetCourierMetaInfo(context.Background(), id, start, end)
	return c.JSON(http.StatusOK, result)
}

func getStartEnd(c echo.Context) (time.Time, time.Time, error) {
	startStr := c.QueryParam(start)
	endStr := c.QueryParam(end)

	start, _ := time.Parse("2006-01-02", startStr)
	//if err != nil {
	//	return time.Time{}, time.Time{}, err
	//}

	end, _ := time.Parse("2006-01-02", endStr)
	//if err != nil {
	//	return time.Time{}, time.Time{}, err
	//}

	return start, end, nil
}
