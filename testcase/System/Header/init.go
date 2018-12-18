package Header

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestHeader() {
	testframework.TFramework.RegTestCase("TestGetIndex", TestGetIndex)
	testframework.TFramework.RegTestCase("TestGetHash", TestGetHash)
	testframework.TFramework.RegTestCase("TestGetPrevHash", TestGetPrevHash)
	testframework.TFramework.RegTestCase("TestGetTimestamp", TestGetTimestamp)
}
