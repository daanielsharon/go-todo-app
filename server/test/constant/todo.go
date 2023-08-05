package constant_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gin-gonic/gin"
)

func TodoCreate(wg *sync.WaitGroup, router *gin.Engine, requestBody []byte, cookie string) (interface{}, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()

	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))
		router.ServeHTTP(recorder, request)
	}()

	wg.Wait()

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	if response.StatusCode == 200 {
		return responseBody["data"], nil
	}

	return nil, errors.New(responseBody["data"].(string))
}

func TodoGet(wg *sync.WaitGroup, router *gin.Engine, cookie string) (interface{}, error) {
	wg.Add(1)

	recorder := httptest.NewRecorder()

	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", Username), nil)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))
		router.ServeHTTP(recorder, request)
	}()

	wg.Wait()

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	if response.StatusCode == 200 {
		return responseBody["data"], nil
	}

	return nil, errors.New(responseBody["data"].(string))
}

func GetTargetId(currentId int64) *int64 {
	var targetId int64
	if currentId == 0 {
		targetId = currentId + 1
	} else {
		targetId = currentId - 1
	}

	return &targetId
}
