package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"server/model/web"
	"server/test/setup"
	"strconv"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// group_id formula, user_id * 3 - 2 (todo), user_id * 3 - 1 (in progress), user_id * 3 (done)

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

	return responseBody["data"], nil
}

func TestCreateTodoSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	res, err := Register("x", &wg, router)

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := Login("x", &wg, router)

	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: groupID,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

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
	assert.Equal(t, "sleep", responseBody["data"].(map[string]interface{})["name"].(string))
}

func TestCreateTodoFailBadRequest(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	res, err := Register("x", &wg, router)

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))

	if err != nil {
		t.FailNow()
	}

	cookie, err := Login("x", &wg, router)

	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		UserID: idInt,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

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

func TestCreateTodoFailUnauthorized(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	requestBody, err := json.Marshal(web.TodoCreateRequest{})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"].(string))
}

func TestCreateTodoFailInternalServerError(t *testing.T) {
	router, db := setup.All()
	defer db.Close()

	var wg sync.WaitGroup

	res, err := Register("x", &wg, router)

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	// nonexistent
	groupID := idInt * 4

	if err != nil {
		t.FailNow()
	}

	cookie, err := Login("x", &wg, router)

	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: groupID,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

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

func TestGetTodoSuccess(t *testing.T) {
	router, db := setup.All()
	defer db.Close()
	username := "x"

	var wg sync.WaitGroup

	res, err := Register(username, &wg, router)

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := Login(username, &wg, router)

	if err != nil {
		t.FailNow()
	}

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	_, err = TodoCreate(&wg, router, requestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", username), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

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

	fmt.Println("responseBody", responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"].(string))
	assert.Equal(t, "sleep", responseBody["data"].([]interface{})[0].(map[string]interface{})["item"].([]interface{})[0].(map[string]interface{})["name"].(string))
}

func TestGetTodoFailUnauthorized(t *testing.T) {
	router, db := setup.All()
	defer db.Close()
	username := "x"

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", username), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println("responseBody", responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"].(string))
}

func TestGetTodoFailNotFound(t *testing.T) {
	router, db := setup.All()
	defer db.Close()
	username := "x"

	var wg sync.WaitGroup

	res, err := Register(username, &wg, router)

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := Login(username, &wg, router)

	if err != nil {
		t.FailNow()
	}

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	_, err = TodoCreate(&wg, router, requestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	// no user with username y registered
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", "y"), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println("responseBody", responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"].(string))
}

func TestGetTodoFailNotFoundUnregistered(t *testing.T) {
	router, db := setup.All()
	defer db.Close()
	username := "x"

	var wg sync.WaitGroup

	_, err := Register(username, &wg, router)

	if err != nil {
		t.FailNow()
	}

	if err != nil {
		t.FailNow()
	}

	cookie, err := Login(username, &wg, router)

	if err != nil {
		t.FailNow()
	}

	wg.Wait()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", 2), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println("responseBody", responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"].(string))
}
