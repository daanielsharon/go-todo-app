package constant_test

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	Username = "x"
	Password = "xx"
)

func RequestBody() *strings.Reader {
	var requestBody = strings.NewReader(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, Username, Password))
	return requestBody
}

func RandomRequestBody() (*strings.Reader, *string, *string) {
	var username string = fmt.Sprintf(`%v%v`, time.Now().Nanosecond(), rand.Intn(1000))
	var password string = fmt.Sprintf(`%v`, rand.Intn(1000))
	var requestBody = strings.NewReader(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, username, password))
	return requestBody, &username, &password
}
