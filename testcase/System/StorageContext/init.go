package StorageContext

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestStorageContext() {
	testframework.TFramework.RegTestCase("TestAsReadOnly", TestAsReadOnly)
}
