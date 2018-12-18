package Datatype

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestDataType() {
	testframework.TFramework.RegTestCase("TestArray", TestArray)
	testframework.TFramework.RegTestCase("TestBoolean", TestBoolean)
	testframework.TFramework.RegTestCase("TestByteArray", TestByteArray)
	testframework.TFramework.RegTestCase("TestInteger", TestInteger)
	testframework.TFramework.RegTestCase("TestReturnType", TestReturnType)
	testframework.TFramework.RegTestCase("TestString", TestString)
}