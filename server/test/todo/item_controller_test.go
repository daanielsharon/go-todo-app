package todotest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"server/model/web"
	constant_test "server/test/constant"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// group_id formula, user_id * 3 - 2 (todo), user_id * 3 - 1 (in progress), user_id * 3 (done)

func TestCreateTodoSuccess(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	testSetup.Wait().Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: groupID,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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
	assert.Equal(t, "sleep", responseBody["data"].(map[string]interface{})["name"].(string))
}

func TestCreateTodoFailBadRequest(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)

	testSetup.Wait().Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		UserID: idInt,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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

func TestCreateTodoFailUnauthorized(t *testing.T) {
	t.Parallel()

	requestBody, err := json.Marshal(web.TodoCreateRequest{})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)

	testSetup.Wait().Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		Name:   "sleep",
		UserID: idInt,
		// nonexistent
		GroupID: time.Now().Nanosecond(),
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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

func TestUpdateTodoSuccess(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	createRequestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	res, err := constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), createRequestBody, loginCookie)
	if err != nil {
		t.FailNow()
	}

	todoId := res.(map[string]interface{})["id"].(float64)
	testSetup.Wait().Wait()

	newGroupId := idInt*3 - 1

	updateRequestBody, err := json.Marshal(web.TodoUpdateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: newGroupId,
	})

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", todoId), bytes.NewReader(updateRequestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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
	assert.Equal(t, "sleep", responseBody["data"].(map[string]interface{})["name"].(string))

	// original position
	assert.Equal(t, groupID, int(res.(map[string]interface{})["groupId"].(float64)))
	assert.Equal(t, newGroupId, int(responseBody["data"].(map[string]interface{})["groupId"].(float64)))

	// fmt.Println("original", int(res.(map[string]interface{})["group_id"].(float64)))
	// fmt.Println("new", int(responseBody["data"].(map[string]interface{})["group_id"].(float64)))
}

func TestUpdateTodoFailedBadRequest(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	createRequestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	_, err = constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), createRequestBody, loginCookie)
	if err != nil {
		t.FailNow()
	}

	testSetup.Wait().Wait()

	newGroupId := idInt*3 - 1

	updateRequestBody, err := json.Marshal(web.TodoUpdateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: newGroupId,
	})

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/todo/0", bytes.NewReader(updateRequestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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

func TestUpdateTodoFailedUnauthorized(t *testing.T) {
	t.Parallel()

	requestBody, err := json.Marshal(web.TodoUpdateRequest{})
	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/todo/1", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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

func TestUpdateTodoFailedNotFound(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)

	testSetup.Wait().Wait()

	newGroupId := idInt*3 - 1
	updateRequestBody, err := json.Marshal(web.TodoUpdateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: newGroupId,
	})

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", int64(time.Now().Nanosecond())), bytes.NewReader(updateRequestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"].(string))
}

func TestGetTodoSuccess(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	_, err = constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), requestBody, loginCookie)
	if err != nil {
		t.FailNow()
	}

	testSetup.Wait().Wait()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", constant_test.Username), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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
	assert.NotNil(t, responseBody["data"].([]interface{})[0])
}

func TestGetTodoFailUnauthorized(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", constant_test.Username), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	_, err = constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), requestBody, loginCookie)
	constant_test.FailIfError(err, t)
	testSetup.Wait().Wait()

	// no user with username y registered
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", "y"), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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
	t.Parallel()

	testSetup.Wait().Wait()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", 2), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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

func TestDeleteTodoSuccess(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	res, err := constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), requestBody, loginCookie)
	if err != nil {
		t.FailNow()
	}

	todoId := res.(map[string]interface{})["id"].(float64)

	testSetup.Wait().Wait()

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", todoId), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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

	fmt.Println("responseBody", responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"].(string))
}

func TestDeleteTodoFailUnauthorized(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", 3), nil)
	fmt.Println("request", request)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

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

func TestDeleteTodoFailBadRequest(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	_, err = constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), requestBody, loginCookie)
	if err != nil {
		t.FailNow()
	}

	// bad request, should be number, got string
	todoId := "sfdkls"

	testSetup.Wait().Wait()

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", todoId), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

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

func TestDeleteTodoFailNotFound(t *testing.T) {
	t.Parallel()

	id := registerResponse.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	constant_test.FailIfError(err, t)
	groupID := idInt*3 - 2

	requestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	res, err := constant_test.TodoCreate(testSetup.Wait(), testSetup.Router(), requestBody, loginCookie)
	constant_test.FailIfError(err, t)

	todoId := res.(map[string]interface{})["id"].(float64)

	testSetup.Wait().Wait()

	// not found, 2%v will never match
	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/2%v", todoId), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", loginCookie))

	recorder := httptest.NewRecorder()
	testSetup.Router().ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"].(string))
}
