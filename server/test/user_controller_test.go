package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	constant_test "server/test/constant"
	"server/test/setup"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", constant_test.RequestBody())
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

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", nil)
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

func TestLoginSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	_, err := constant_test.Register(&wg, router)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", constant_test.RequestBody())
	fmt.Println("requestBody(", constant_test.RequestBody())
	fmt.Println("request", request)
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
	fmt.Println("cookie", cookie)

	assert.Contains(t, cookie, "token")
}

func TestLoginFailBadRequest(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	_, err := constant_test.Register(&wg, router)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody := strings.NewReader(`{"username":"", "password":""}`)
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

func TestLogoutSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	_, err := constant_test.Register(&wg, router)
	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(&wg, router)

	wg.Wait()

	assert.NotEqual(t, nil, cookie)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

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
	newCookie := headers.Get("Set-Cookie")

	assert.NotContains(t, newCookie, cookie)
}
