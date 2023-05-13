package main

//
//import (
//	"github.com/golang/mock/gomock"
//	"net/http"
//	"net/http/httptest"
//	mock_courier "sss/repository/courier/mocks"
//	mock_order "sss/repository/order/mocks"
//	"sync"
//	"testing"
//	"time"
//)
//
//func TestApp(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	repoC := mock_courier.NewMockRepo(ctl)
//	repoO := mock_order.NewMockRepo(ctl)
//
//	echo := app(repoO, repoC)
//
//	ts := httptest.NewServer(echo)
//	defer ts.Close()
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
//				else {
//					if resp.StatusCode != 200 {
//						t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
//					}
//				}
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
//}
