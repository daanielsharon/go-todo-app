package constant_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func Register(wg *sync.WaitGroup, router *gin.Engine) (interface{}, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", RequestBody())
		request.Header.Add("Content-Type", "application/json")
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

	return responseBody["data"], nil
}

func RandomRegister(wg *sync.WaitGroup, router *gin.Engine) (*string, *string, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()
	requestBody, username, password := RandomRequestBody()
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", requestBody)
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)
	}()

	wg.Wait()

	_ = recorder.Result()

	return username, password, nil
}

func Login(wg *sync.WaitGroup, router *gin.Engine) (string, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", RequestBody())
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)
	}()

	wg.Wait()

	response := recorder.Result()
	headers := response.Header
	cookie := headers.Get("Set-Cookie")

	if cookie != "" {
		return strings.Split(cookie, "=")[1], nil
	}

	err := fmt.Errorf("No cookie obtaind")
	return "", err
}

func CustomLogin(wg *sync.WaitGroup, router *gin.Engine, username *string, password *string) (string, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()
	auth := strings.NewReader(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, *username, *password))
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", auth)
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)
	}()

	wg.Wait()

	response := recorder.Result()
	headers := response.Header
	cookie := headers.Get("Set-Cookie")

	if cookie != "" {
		return strings.Split(cookie, "=")[1], nil
	}

	err := fmt.Errorf("No cookie obtaind")
	return "", err
}
