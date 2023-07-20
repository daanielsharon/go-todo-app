package test

import (
	"bytes"
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
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

// group_id formula, user_id * 3 - 2 (todo), user_id * 3 - 1 (in progress), user_id * 3 (done)

func TodoCreate(wg *sync.WaitGroup, router *gin.Engine, requestBody []byte) (interface{}, error) {
	wg.Add(1)
	recorder := httptest.NewRecorder()

	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/", bytes.NewReader(requestBody))
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
}
