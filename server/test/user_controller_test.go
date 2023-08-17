package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	constant_test "server/test/constant"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	t.Parallel()

	requestBody, username, _ := constant_test.RandomRequestBody()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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
	assert.Equal(t, *username, responseBody["data"].(map[string]interface{})["username"])
}

func TestRegisterFailBadRequest(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/register", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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
	t.Parallel()

	username, password, err := constant_test.RandomRegister(testSetup.Wait(), testSetup.Router())
	constant_test.FailIfError(err, t)
	auth := strings.NewReader(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, *username, *password))

	testSetup.Wait().Wait()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", auth)
	request.Header.Add("Content-Type", "application/json")

	testSetup.Router().ServeHTTP(recorder, request)
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
	t.Parallel()

	_, _, err := constant_test.RandomRegister(testSetup.Wait(), testSetup.Router())
	constant_test.FailIfError(err, t)

	testSetup.Wait().Wait()

	requestBody := strings.NewReader(`{"username":"", "password":""}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/login", requestBody)
	request.Header.Add("Content-Type", "application/json")

	testSetup.Router().ServeHTTP(recorder, request)
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
	t.Parallel()

	username, password, err := constant_test.RandomRegister(testSetup.Wait(), testSetup.Router())
	constant_test.FailIfError(err, t)

	cookie, err := constant_test.CustomLogin(testSetup.Wait(), testSetup.Router(), username, password)

	testSetup.Wait().Wait()

	assert.NotEqual(t, nil, cookie)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/users/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	testSetup.Router().ServeHTTP(recorder, request)
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
