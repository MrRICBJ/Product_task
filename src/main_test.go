package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"sync"
	"testing"
	"time"
)

//func TestApp(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	repoC := mock_courier.NewMockRepo(ctl)
//	repoO := mock_order.NewMockRepo(ctl)
//	ctx := context.Background()
//
//	mockRepo := []entity.Courier{
//		{
//			CourierId:    1,
//			CourierType:  "FOOT",
//			Regions:      []int32{2, 4},
//			WorkingHours: []string{"12:34:00", "12:12:00"},
//		},
//	}
//
//	//expected := &dto.GetCouriersResponse{
//	//	Couriers: []dto.CourierDto{
//	//		{
//	//			CourierId:    1,
//	//			CourierType:  "FOOT",
//	//			Regions:      []int32{2, 4},
//	//			WorkingHours: []string{"12:34:00", "12:12:00"},
//	//		},
//	//	},
//	//	Limit:  int32(1),
//	//	Offset: int32(0),
//	//}
//
//	repoC.EXPECT().GetAll(ctx, int32(1), int32(0)).Return(mockRepo, nil).Times(1)
//	echo := app(repoO, repoC)
//
//	ts := httptest.NewServer(echo)
//	defer ts.Close()
//
//	endpoints := []string{"/couriers"}
//	requestsPerSecond := 15
//	totalRequests := 100
//
//	var wg sync.WaitGroup
//	for _, endpoint := range endpoints {
//		for i := 0; i < totalRequests; i++ {
//			wg.Add(1)
//			go func(endpoint string) {
//				defer wg.Done()
//
//				// Выполняем запрос
//				resp, err := http.Get(ts.URL + endpoint)
//				if err != nil {
//					t.Errorf("Error making request: %v", err)
//				}
//
//				// Проверяем, что сервис возвращает код 429, если количество запросов превышает лимит
//				if i >= requestsPerSecond*10 {
//					if resp.StatusCode != 429 {
//						t.Errorf("Expected status code 429, but got %d", resp.StatusCode)
//					}
//				}
//
//				// Проверяем, что каждый эндпоинт не превышает лимит в 10 RPS
//				//else {
//				//	if resp.StatusCode != 200 {
//				//		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
//				//	}
//				//}
//			}(endpoint)
//
//			// Ожидаем некоторое время между запросами, чтобы не превышать лимит в 10 RPS
//			time.Sleep(time.Second / time.Duration(requestsPerSecond))
//		}
//	}
//
//	// Ожидаем завершения всех запросов
//	wg.Wait()
//
//	//require.Error(t, t.Erro)
//}

func TestRateLimiter(t *testing.T) {
	e := echo.New()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	go func() {
		if err := e.Start(":8080"); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			resp, err := http.Get("http://localhost:8080/")
			if err != nil {
				fmt.Println("Request error:", err)
			} else {
				fmt.Println("Response status:", resp.Status)
			}
			wg.Done()
		}()
		wg.Wait()
		if err := e.Shutdown(context.Background()); err != nil {
			fmt.Println("Server shutdown error:", err)
		}
	}
}
