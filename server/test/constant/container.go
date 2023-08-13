package constant_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gin-gonic/gin"
)

func ContainerAdd(wg *sync.WaitGroup, router *gin.Engine, cookie string, requestBody []byte) {
	fmt.Println("executed first")
	wg.Add(1)

	recorder := httptest.NewRecorder()
	go func() {
		defer wg.Done()
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/todo/container/", bytes.NewReader(requestBody))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Cookie", fmt.Sprintf("token=%v", cookie))
		router.ServeHTTP(recorder, request)
	}()

	wg.Wait()
}
