package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"server/test/setup"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Register(request string, wg *sync.WaitGroup, router *gin.Engine) (interface{}, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(fmt.Sprintf(`{"username":"%v"}`, request))
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", requestBody)
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

func Login(request string, wg *sync.WaitGroup, router *gin.Engine) (string, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(fmt.Sprintf(`{"username":"%v"}`, request))
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", requestBody)
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

func TestRegisterSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	requestBody := strings.NewReader(`{"username":"x"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"].(string))
	assert.Equal(t, "x", responseBody["data"].(map[string]interface{})["username"])
}

func TestRegisterFailBadRequest(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	requestBody := strings.NewReader(`{"username":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"].(string))
}

func TestRegisterFailInternalServerError(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 500, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 500, int(responseBody["code"].(float64)))
	assert.Equal(t, "Internal Server Error", responseBody["status"].(string))
}

func TestLoginSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	_, err := Register("x", &wg, router)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody := strings.NewReader(`{"username":"x"}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	headers := response.Header
	cookie := headers.Get("Set-Cookie")

	assert.Contains(t, cookie, "token")
}

func TestLoginFailBadRequest(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	_, err := Register("x", &wg, router)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody := strings.NewReader(`{"username":""}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])

	headers := response.Header
	cookie := headers.Get("Set-Cookie")

	assert.NotContains(t, cookie, "token")
}

func TestLoginFailInternalServerError(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	_, err := Register("x", &wg, router)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody := strings.NewReader(``)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 500, int(responseBody["code"].(float64)))
	assert.Equal(t, "Internal Server Error", responseBody["status"])

	headers := response.Header
	cookie := headers.Get("Set-Cookie")

	assert.NotContains(t, cookie, "token")
}
