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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePrioritySuccess(t *testing.T) {
	setup := setup.NewTestSetup()
	setup.Open()
	defer setup.Close()

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

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/priority/%v", originId), bytes.NewReader(requestBody))
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

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/priority/%v", originId), bytes.NewReader(requestBody))
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

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/priority/%v", originId), bytes.NewReader(requestBody))
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

	request := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("http://localhost:8080/api/v1/todo/priority/%v", 0), bytes.NewReader(requestBody))
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
