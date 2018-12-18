package Operator

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestNeoVMOperator() {
	testframework.TFramework.RegTestCase("TestOperationAdd", TestOperationAdd)
	testframework.TFramework.RegTestCase("TestOperationAnd", TestOperationAnd)
	testframework.TFramework.RegTestCase("TestOperationDivide", TestOperationDivide)
	testframework.TFramework.RegTestCase("TestOperationEqual", TestOperationEqual)
	testframework.TFramework.RegTestCase("TestOperationLargerEqual", TestOperationLargerEqual)
	testframework.TFramework.RegTestCase("TestOperationLeftShift", TestOperationLeftShift)
	testframework.TFramework.RegTestCase("TestOperationLarger", TestOperationLarger)
	testframework.TFramework.RegTestCase("TestOperationMode", TestOperationMode)
	testframework.TFramework.RegTestCase("TestOperationMulti", TestOperationMulti)
	testframework.TFramework.RegTestCase("TestOperationNotEqual", TestOperationNotEqual)
	testframework.TFramework.RegTestCase("TestOperationNegative", TestOperationNegative)
	testframework.TFramework.RegTestCase("TestOperationOr", TestOperationOr)
	testframework.TFramework.RegTestCase("TestOperationRightShift", TestOperationRightShift)
	testframework.TFramework.RegTestCase("TestOperationSmallerEqual", TestOperationSmallerEqual)
	testframework.TFramework.RegTestCase("TestOperationSelfAdd", TestOperationSelfAdd)
	testframework.TFramework.RegTestCase("TestOperationSelfSub", TestOperationSelfSub)
	testframework.TFramework.RegTestCase("TestOperationSmaller", TestOperationSmaller)
	testframework.TFramework.RegTestCase("TestOperationSub", TestOperationSub)
}

