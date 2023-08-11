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

func TestUpdatePrioritySuccess(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	_, err := constant_test.Register(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	cookie, err := constant_test.Login(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	todo, err := constant_test.TodoGet(setup.Wait(), setup.Router(), cookie)

	originId := todo.([]interface{})[0].(map[string]interface{})["id"].(float64)
	originPriority := todo.([]interface{})[0].(map[string]interface{})["priority"].(float64)

	targetId := todo.([]interface{})[2].(map[string]interface{})["id"].(float64)
	targetPriority := todo.([]interface{})[2].(map[string]interface{})["priority"].(float64)

	requestBody, err := json.Marshal(web.TodoUpdatePriority{
		OriginPriority: int64(targetPriority),
		TargetID:       int64(targetId),
		TargetPriority: int64(originPriority),
	})
	constant_test.FailIfError(err, t)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/container/priority/%v", originId), bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Result().StatusCode)

	newTodo, err := constant_test.TodoGet(setup.Wait(), setup.Router(), cookie)
	data := newTodo.([]interface{})

	newOriginId := data[0].(map[string]interface{})["id"].(float64)
	newTargetId := data[2].(map[string]interface{})["id"].(float64)

	// both id and priority get swapped
	// but it will be sorted by priority
	// the first should be the last now
	assert.Equal(t, originId, newTargetId)
	assert.Equal(t, targetId, newOriginId)
}

func TestUpdatePriorityBadRequest(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	_, err := constant_test.Register(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	cookie, err := constant_test.Login(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	todo, err := constant_test.TodoGet(setup.Wait(), setup.Router(), cookie)
	originId := todo.([]interface{})[0].(map[string]interface{})["id"].(float64)

	requestBody, err := json.Marshal(web.TodoUpdatePriority{
		OriginPriority: 0,
		TargetID:       0,
		TargetPriority: 0,
	})
	constant_test.FailIfError(err, t)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/container/priority/%v", originId), bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	constant_test.FailIfError(err, t)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"].(string))

}
func TestUpdatePriorityUnauthorized(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	_, err := constant_test.Register(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	cookie, err := constant_test.Login(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	todo, err := constant_test.TodoGet(setup.Wait(), setup.Router(), cookie)
	originId := todo.([]interface{})[0].(map[string]interface{})["id"].(float64)

	requestBody, err := json.Marshal(web.TodoUpdatePriority{
		OriginPriority: 0,
		TargetID:       0,
		TargetPriority: 0,
	})
	constant_test.FailIfError(err, t)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/container/priority/%v", originId), bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	constant_test.FailIfError(err, t)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"].(string))
}
func TestUpdatePriorityNotFound(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

	t.Parallel()

	_, err := constant_test.Register(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)
	_, err = constant_test.Login(setup.Wait(), setup.Router())
	constant_test.FailIfError(err, t)

	requestBody, err := json.Marshal(web.TodoUpdatePriority{
		OriginPriority: 0,
		TargetID:       0,
		TargetPriority: 0,
	})
	constant_test.FailIfError(err, t)

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/container/priority/%v", 0), bytes.NewReader(requestBody))
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	setup.Router().ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	constant_test.FailIfError(err, t)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"].(string))
}

func TestContainerCreate(t *testing.T) {
	t.Run("create container should succeed", func(t *testing.T) {
		t.Parallel()

		setup := setup.NewTestSetup()
		setup.Open()
		defer setup.Close()

		res, err := constant_test.Register(setup.Wait(), setup.Router())
		constant_test.FailIfError(err, t)
		cookie, err := constant_test.Login(setup.Wait(), setup.Router())
		constant_test.FailIfError(err, t)

		id := res.(map[string]interface{})["id"]
		idInt, err := strconv.Atoi(fmt.Sprint(id))

		fmt.Println("idINt", idInt)

		requestBody, err := json.Marshal(web.ContainerCreateRequest{
			UserId:    int64(idInt),
			GroupName: "pending",
			Priority:  4,
		})

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/container/", bytes.NewReader(requestBody))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

		recorder := httptest.NewRecorder()
		setup.Router().ServeHTTP(recorder, request)

		response := recorder.Result()
		assert.Equal(t, 200, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		constant_test.FailIfError(err, t)

		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 200, int(responseBody["code"].(float64)))
		assert.Equal(t, "OK", responseBody["status"].(string))
		assert.Equal(t, "pending", responseBody["data"].(map[string]interface{})["groupName"].(string))

		todo, err := constant_test.TodoGet(setup.Wait(), setup.Router(), cookie)
		assert.Equal(t, "pending", todo.([]interface{})[3].(map[string]interface{})["group_name"].(string))
	})

	t.Run("create container should be unauthorized", func(t *testing.T) {
		t.Parallel()

		setup := setup.NewTestSetup()
		setup.Open()
		defer setup.Close()

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/container/", nil)
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		setup.Router().ServeHTTP(recorder, request)

		response := recorder.Result()
		assert.Equal(t, 401, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		constant_test.FailIfError(err, t)

		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 401, int(responseBody["code"].(float64)))
		assert.Equal(t, "Unauthorized", responseBody["status"].(string))
	})
	t.Run("create container with wrong request body", func(t *testing.T) {
		t.Parallel()

		setup := setup.NewTestSetup()
		setup.Open()
		defer setup.Close()

		res, err := constant_test.Register(setup.Wait(), setup.Router())
		constant_test.FailIfError(err, t)
		cookie, err := constant_test.Login(setup.Wait(), setup.Router())
		constant_test.FailIfError(err, t)

		id := res.(map[string]interface{})["id"]
		idInt, err := strconv.Atoi(fmt.Sprint(id))

		requestBody, err := json.Marshal(web.ContainerCreateRequest{
			UserId: int64(idInt),
		})

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/container/", bytes.NewReader(requestBody))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

		recorder := httptest.NewRecorder()
		setup.Router().ServeHTTP(recorder, request)

		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		constant_test.FailIfError(err, t)

		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request", responseBody["status"].(string))
	})
	t.Run("create container exceeding limit (5) in container should be bad request", func(t *testing.T) {
		t.Parallel()

		setup := setup.NewTestSetup()
		setup.Open()
		defer setup.Close()

		res, err := constant_test.Register(setup.Wait(), setup.Router())
		constant_test.FailIfError(err, t)
		cookie, err := constant_test.Login(setup.Wait(), setup.Router())
		constant_test.FailIfError(err, t)

		id := res.(map[string]interface{})["id"]
		idInt, err := strconv.Atoi(fmt.Sprint(id))

		requestBody, err := json.Marshal(web.ContainerCreateRequest{
			UserId:    int64(idInt),
			GroupName: "pending",
			Priority:  4,
		})

		for i := 0; i <= 2; i++ {
			constant_test.ContainerAdd(setup.Wait(), setup.Router(), cookie, requestBody)
		}

		setup.Wait().Wait()

		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/container/", bytes.NewReader(requestBody))

		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))

		recorder := httptest.NewRecorder()
		setup.Router().ServeHTTP(recorder, request)

		response := recorder.Result()
		assert.Equal(t, 400, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		constant_test.FailIfError(err, t)

		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "Bad Request", responseBody["status"].(string))
		assert.Equal(t, "cannot create more than 5 containers!", responseBody["data"].(string))
	})
}
