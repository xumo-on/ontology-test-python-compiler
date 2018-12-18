package Native

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestNative() {
	testframework.TFramework.RegTestCase("TestInvoke", TestInvoke)
}
