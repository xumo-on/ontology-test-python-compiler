package Contract

import "github.com/xumo-on/ontology-test-python-compiler/testframework"

func TestContract() {
	testframework.TFramework.RegTestCase("TestGetScript", TestGetScript)
	testframework.TFramework.RegTestCase("TestMigrate", TestMigrate)
}
