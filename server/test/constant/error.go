package constant_test

import "testing"

func FailIfError(err error, t *testing.T) {
	if err != nil {
		t.FailNow()
	}
}
