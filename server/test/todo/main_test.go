package todotest

import (
	"os"
	"server/helper"
	constant_test "server/test/constant"
	"server/test/setup"
	"testing"
)

var (
	registerResponse interface{}
	loginCookie      string
	testSetup        setup.TestSetup
)

func TestMain(m *testing.M) {
	// setup := setup.NewTestSetup()
	// setup.Open()

	// res, err := constant_test.Register(setup.Wait(), setup.Router())
	// helper.PanicIfError(err)
	// registerResponse = res

	// cookie, err := constant_test.Login(setup.Wait(), setup.Router())
	// helper.PanicIfError(err)
	// loginCookie = cookie

	// testSetup = setup

	// m.Run()

	// testSetup.TruncateAll()
	// testSetup.Close()

	os.Exit(testMainExitWrapper(m))
}

func testMainExitWrapper(m *testing.M) int {
	setup := setup.NewTestSetup()
	setup.Open()

	res, err := constant_test.Register(setup.Wait(), setup.Router())
	helper.PanicIfError(err)
	registerResponse = res

	cookie, err := constant_test.Login(setup.Wait(), setup.Router())
	helper.PanicIfError(err)
	loginCookie = cookie

	testSetup = setup

	defer testSetup.TruncateAll()
	defer testSetup.Close()

	return m.Run()
}
