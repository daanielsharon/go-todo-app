package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"server/test/setup"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		t.Fail()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"].(string))
	assert.Equal(t, "x", responseBody["data"].(map[string]interface{})["username"])

	headers := response.Header
	cookie := headers.Get("Set-Cookie")

	assert.NotContains(t, cookie, "token")
}

// func TestRegisterFail(t *testing.T)
