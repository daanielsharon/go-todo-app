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
	"server/test/setup"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// group_id formula, user_id * 3 - 2 (todo), user_id * 3 - 1 (in progress), user_id * 3 (done)

func TestCreateTodoSuccess(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: groupID,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		UserID: idInt,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	requestBody, err := json.Marshal(web.TodoCreateRequest{})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

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

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	requestBody, err := json.Marshal(web.TodoCreateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: groupID,
	})

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	createRequestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	res, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), createRequestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	todoId := res.(map[string]interface{})["id"].(float64)
	setup.Wait().Wait()

	newGroupId := idInt*3 - 1

	updateRequestBody, err := json.Marshal(web.TodoUpdateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: newGroupId,
	})

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", todoId), bytes.NewReader(updateRequestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	assert.Equal(t, "sleep", responseBody["data"].(map[string]interface{})["name"].(string))

	// original position
	assert.Equal(t, groupID, int(res.(map[string]interface{})["groupId"].(float64)))
	assert.Equal(t, newGroupId, int(responseBody["data"].(map[string]interface{})["groupId"].(float64)))

	// fmt.Println("original", int(res.(map[string]interface{})["group_id"].(float64)))
	// fmt.Println("new", int(responseBody["data"].(map[string]interface{})["group_id"].(float64)))
}

func TestUpdateTodoFailedBadRequest(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	createRequestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	res, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), createRequestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	newGroupId := idInt*3 - 1

	updateRequestBody, err := json.Marshal(web.TodoUpdateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: newGroupId,
	})

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/todo/0", bytes.NewReader(updateRequestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	requestBody, err := json.Marshal(web.TodoUpdateRequest{})
	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8080/api/v1/todo/1", bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	createRequestBody, err := json.Marshal(
		web.TodoCreateRequest{
			Name:    "sleep",
			UserID:  idInt,
			GroupID: groupID,
		},
	)

	res, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), createRequestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	todoId := res.(map[string]interface{})["id"].(float64)

	setup.Wait().Wait()

	newGroupId := idInt*3 - 1
	updateRequestBody, err := json.Marshal(web.TodoUpdateRequest{
		Name:    "sleep",
		UserID:  idInt,
		GroupID: newGroupId,
	})

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", int64(todoId)+2), bytes.NewReader(updateRequestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

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

	_, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), requestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", constant_test.Username), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	assert.Equal(t, "sleep", responseBody["data"].([]interface{})[0].(map[string]interface{})["item"].([]interface{})[0].(map[string]interface{})["name"].(string))
}

func TestGetTodoFailUnauthorized(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	username := "x"

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", username), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

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

	_, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), requestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	// no user with username y registered
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", "y"), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	_, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	setup.Wait().Wait()

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", 2), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

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

	res, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), requestBody, cookie)
	if err != nil {
		t.FailNow()
	}
	fmt.Println("res todo", res)

	todoId := res.(map[string]interface{})["id"].(float64)

	setup.Wait().Wait()

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", todoId), nil)
	fmt.Println("request", request)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", 3), nil)
	fmt.Println("request", request)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

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

	_, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), requestBody, cookie)
	if err != nil {
		t.FailNow()
	}

	// bad request, should be number, got string
	todoId := "sfdkls"

	setup.Wait().Wait()

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/%v", todoId), nil)
	fmt.Println("request", request)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.FailNow()
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println("responseBody", responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"].(string))
}

func TestDeleteTodoFailNotFound(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	res, err := constant_test.Register(setup.Wait(), setup.Router())

	if err != nil {
		t.FailNow()
	}

	id := res.(map[string]interface{})["id"]
	idInt, err := strconv.Atoi(fmt.Sprint(id))
	groupID := idInt*3 - 2

	if err != nil {
		t.FailNow()
	}

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())

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

	res, err = constant_test.TodoCreate(setup.Wait(), setup.Router(), requestBody, cookie)
	if err != nil {
		t.FailNow()
	}
	fmt.Println("res todo", res)

	todoId := res.(map[string]interface{})["id"].(float64)

	setup.Wait().Wait()

	// not found, 2%v will never match
	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/todo/2%v", todoId), nil)
	fmt.Println("request", request)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

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
