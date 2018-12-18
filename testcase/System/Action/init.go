package Action

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestAction() {
	testframework.TFramework.RegTestCase("TestRegisterAction", TestRegisterAction)
}
