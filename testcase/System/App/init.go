package App

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestApp() {
	testframework.TFramework.RegTestCase("TestRegisterAppCall", TestRegisterAppCall)
}
