package constant_test

import (
	"fmt"
	"strings"
)

const (
	Username = "x"
	Password = "xx"
)

func RequestBody() *strings.Reader {
	var requestBody = strings.NewReader(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, Username, Password))
	return requestBody
}
